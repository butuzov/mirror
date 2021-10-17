package mirror

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"os"
	"sort"
	"strings"

	"github.com/butuzov/mirror/internal/checker"
	"github.com/butuzov/mirror/internal/imports"
	"github.com/butuzov/mirror/internal/rules"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type analyzer struct {
	imports  map[string][]imports.KV // Alias & name of imports sorted per file
	checkers []*checker.Checker      // Available checkers.
}

func NewAnalyzer() *analysis.Analyzer {
	a := analyzer{
		imports: map[string][]imports.KV{},
		checkers: []*checker.Checker{
			rules.NewRegexpChecker(),
			rules.NewStringsChecker(),
			rules.NewBytesChecker(),
		},
	}

	return &analysis.Analyzer{
		Name: "mirror",
		Doc:  "looks for mirror patterns",
		Run:  a.run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
		Flags: flags(),
	}
}

func flags() flag.FlagSet {
	set := flag.NewFlagSet("", flag.PanicOnError)
	set.Bool("with-tests", false, "do not skip tests in reports")
	set.Bool("with-debug", false, "debug linter run (development only)")
	return *set
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	// read flags...
	withDebug := pass.Analyzer.Flags.Lookup("with-debug").Value.String() == "true"
	withTests := pass.Analyzer.Flags.Lookup("with-tests").Value.String() == "true"

	ins, _ := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	a.loadImports(pass.Fset, ins)

	// checking phase
	issues := []*analysis.Diagnostic{}

	ins.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
		call, _ := n.(*ast.CallExpr)

		debugFn := debugNoOp
		if withDebug {
			debugFn = debug(pass.Fset, pass.TypesInfo.Types)
		}

		for _, check := range a.checkers {
			// refresh debug function with updated fset and types.
			imports := a.locaImports(call.Pos(), pass.Fset)

			// setup checker
			check.
				WithDebug(debugFn).
				WithTypes(pass.TypesInfo.Types).
				WithImports(imports)

			violation := check.Check(call)
			if violation == nil {
				continue
			}

			issues = append(issues, &analysis.Diagnostic{
				Pos:     n.Pos(),
				End:     n.End(),
				Message: violation.Message,
				// -------- Suggestion -----------------------------------------------
				SuggestedFixes: (suggest{
					fset:    pass.Fset,
					call:    call,
					data:    violation,
					imports: imports,
				}).Get(),
				// -------- Suggestion -----------------------------------------------
			})
		}
	})

	// --- Reporting issues ------------------------------------------------------
	for i := range issues {
		// excluding tests (default action)
		fileIsTest := isTest(pass.Fset.Position(issues[i].Pos).Filename)
		if !withTests && fileIsTest {
			continue
		}
		pass.Report(*issues[i])
	}

	return nil, nil
}

func isTest(fileName string) bool {
	return strings.HasSuffix(fileName, string("_test.go"))
}

func (a *analyzer) loadImports(fs *token.FileSet, ins *inspector.Inspector) {
	ins.Preorder([]ast.Node{(*ast.ImportSpec)(nil)}, func(node ast.Node) {
		is, _ := node.(*ast.ImportSpec)

		// what file we working with?
		filename := fs.Position(node.Pos()).Filename

		var (
			name  = strings.Trim(is.Path.Value, `"`)
			alias = is.Name.String()
		)

		if is.Name == nil {
			alias = name
		}

		a.imports[filename] = append(a.imports[filename], imports.KV{
			Key: name,
			Val: alias,
		})
	})

	for k := range a.imports {
		if len(a.imports[k]) >= 13 {
			k := k
			sort.Slice(a.imports[k], func(i, j int) bool {
				return a.imports[k][i].Val < a.imports[k][j].Val
			})
		}
	}
}

// localImports (it relation to token.Pos) found in fileset
func (a *analyzer) locaImports(p token.Pos, fs *token.FileSet) []imports.KV {
	fileName := fs.Position(p).Filename
	if v, ok := a.imports[fileName]; ok {
		return v
	}
	return nil
}

// --- Suggestions -------------------------------------------------------------

type suggest struct {
	fset    *token.FileSet
	call    *ast.CallExpr
	data    *checker.Violation
	imports []imports.KV
}

func (s suggest) Get() []analysis.SuggestedFix {
	var suggestions []analysis.SuggestedFix

	suggestions = append(suggestions, analysis.SuggestedFix{
		Message: "Replace function",
		TextEdits: []analysis.TextEdit{
			{
				Pos:     s.call.Pos(),
				End:     s.call.Pos(),
				NewText: s.Bytes(),
			},
		},
	})

	// does it need new import?
	if s.data.Type == checker.Method {
		return suggestions
	}

	for i := range s.imports {
		if s.imports[i].Key == s.data.Alternative.Package {
			return suggestions
		}
	}

	// todo(butuzov): add import
	// fmt.Printf("(1) add import for %s\n", s.data.Alternative.Package)

	return suggestions
}

func (s suggest) Bytes() []byte {
	var suggest bytes.Buffer

	switch s.data.Type {
	case checker.Method:
		se, _ := s.call.Fun.(*ast.SelectorExpr)
		receiver, _ := se.X.(*ast.Ident)

		suggest.WriteString(receiver.String())
		suggest.WriteByte('.')
		suggest.WriteString(s.data.Alternative.Method)

	case checker.Function:

		var name string

		switch v := s.call.Fun.(type) {
		case *ast.SelectorExpr:
			// named (regular) import
			receiver, _ := v.X.(*ast.Ident)
			name = s.importsReplacementLookup(receiver.String(), s.data.Alternative.Package)
		case *ast.Ident:
			name = s.importsReplacementLookup(".", s.data.Alternative.Package)
		}

		if name != "." {
			suggest.WriteString(name)
			suggest.WriteByte('.')
		}
		suggest.WriteString(s.data.Alternative.Function)

	default:
		panic("what is that? not implemented")
	}

	suggest.WriteByte('(')

	// arguments.
	var args []string
	for index, v := range s.call.Args {
		if e, ok := s.data.AltArgAt(index); ok {
			args = append(args, types.ExprString(e))
			continue
		}
		args = append(args, types.ExprString(v))
	}

	suggest.WriteString(strings.Join(args, ", "))
	suggest.WriteByte(')')

	return suggest.Bytes()
}

// Scenario, we going to replace bytes with strings
//
//                             | bytes | strings | str | .
//                             ----------------------------
//                       bytes |   x   |    x    |  x  | x
//  current package      b     |   x   |    x    |  x  | x
//                       .     |   x   |    x    |  x  | x
//
func (s suggest) importsReplacementLookup(pkgCur, pkgNew string) string {
	// It's same package (with no alias import).
	if pkgCur == pkgNew {
		return pkgCur
	}

	// Searching for real name of dot imported, aliased or regularly
	// imported package.
	var pkgCurName string
	for i := range s.imports {
		if s.imports[i].Val == pkgCur {
			pkgCurName = s.imports[i].Key
			break
		}
	}

	// If its same as new one - return it.
	if pkgCurName == pkgNew {
		return pkgCur
	}

	// Does the package already imported?
	for i := range s.imports {
		if s.imports[i].Key == pkgNew {
			return s.imports[i].Val
		}
	}

	return pkgNew
}

// --- debug -------------------------------------------------------------------

func debug(f *token.FileSet, t map[ast.Expr]types.TypeAndValue) func(ast.Expr) {
	return func(node ast.Expr) {
		tv := t[node]
		Errorf("how to work with %T of type %#v ?\n", node, tv)

		fmt.Print("affected line> ")
		printer.Fprint(os.Stdout, f, node)
		fmt.Println()
		fmt.Println(strings.Repeat("^", 80))
	}
}

func debugNoOp(_ ast.Expr) {}

// Errorf is function that reports error
const (
	prefixErr  = "\033[91m [ERR] \033[0m"
	prefixWarn = "\033[93m [ERR] \033[0m"
	prefixInfo = "\033[94m [ERR] \033[0m"
)

func Errorf(format string, args ...interface{}) {
	fmt.Printf(prefixErr+format, args...)
}

func Warnf(format string, args ...interface{}) {
	fmt.Printf(prefixWarn+format, args...)
}

func Infof(format string, args ...interface{}) {
	fmt.Printf(prefixInfo+format, args...)
}
