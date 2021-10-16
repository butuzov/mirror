package mirror

import (
	"flag"
	"go/ast"
	"go/token"
	"strings"

	"github.com/butuzov/mirror/internal/checker"
	"github.com/butuzov/mirror/internal/rules"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type analyzer struct {
	imports  map[string]map[string][]string
	checkers []*checker.Checker
}

func NewAnalyzer() *analysis.Analyzer {
	a := analyzer{
		imports: map[string]map[string][]string{},
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
	set.Bool("with-debug", false, "do not skip tests in reports")
	return *set
}

func (a *analyzer) populateImports(fs *token.FileSet, ins *inspector.Inspector) {
	ins.Preorder([]ast.Node{(*ast.ImportSpec)(nil)}, func(node ast.Node) {
		fname := fs.Position(node.Pos()).Filename
		_, ok := a.imports[fname]
		if !ok {
			a.imports[fname] = make(map[string][]string)
		}

		// grab imports.
		is, _ := node.(*ast.ImportSpec)
		key := strings.Trim(is.Path.Value, `"`)
		name := is.Name.String()

		if is.Name == nil {
			name = key
		}

		a.imports[fname][key] = append(a.imports[fname][key], name)
	})
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	withDebug := pass.Analyzer.Flags.Lookup("with-debug").Value.String() == "true"
	withTests := pass.Analyzer.Flags.Lookup("with-tests").Value.String() == "true"

	ins, _ := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// prepare phase
	a.populateImports(pass.Fset, ins)

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
			localImports := a.importsFor(call.Pos(), pass.Fset)

			check.
				WithDebug(debugFn).
				WithTypes(pass.TypesInfo.Types).
				WithImports(localImports)

			violation := check.Check(call)
			if violation == nil {
				continue
			}

			issues = append(issues, &analysis.Diagnostic{
				Pos:     n.Pos(),
				End:     n.End(),
				Message: violation.Message,
				SuggestedFixes: []analysis.SuggestedFix{
					// Creating suggestion.
					(suggest{
						fset:    pass.Fset,
						call:    call,
						data:    violation,
						imports: localImports,
					}).Export(),
				},
			})

		}
	})

	// printing results
	for i := range issues {
		if !withTests && isTest(pass.Fset.PositionFor(issues[i].Pos, true).Filename) {
			continue
		}
		pass.Report(*issues[i])
	}

	return nil, nil
}

func isTest(fileName string) bool {
	return strings.HasSuffix(fileName, string("_test.go"))
}

func PosAtFile(p token.Pos, fs *token.FileSet) string {
	return fs.Position(p).Filename
}

func (a *analyzer) importsFor(p token.Pos, fs *token.FileSet) map[string][]string {
	fileName := fs.Position(p).Filename
	if v, ok := a.imports[fileName]; ok {
		return v
	}
	return nil
}
