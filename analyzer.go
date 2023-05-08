package mirror

import (
	"flag"
	"go/ast"
	"strings"
	"sync"

	"github.com/butuzov/mirror/internal/checker"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type analyzer struct {
	checkers map[string]*checker.Checker // Available checkers.
	imports  map[string]checker.Import   // Map

	withTests bool
	withDebug bool

	once sync.Once
}

func NewAnalyzer() *analysis.Analyzer {
	flags := flags()

	a := analyzer{
		imports:  map[string]checker.Import{},
		checkers: make(map[string]*checker.Checker),
	}

	a.register(newRegexpChecker())
	a.register(newStringsChecker())
	a.register(newBytesChecker())
	a.register(newMaphashChecker())
	a.register(newBufioChecker())
	a.register(newOsChecker())
	a.register(newUTF8Checker())
	a.register(newHTTPTestChecker())

	return &analysis.Analyzer{
		Name: "mirror",
		Doc:  "reports wrong mirror patterns of bytes/strings usage",
		Run:  a.run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
		Flags: flags,
	}
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	issues := []*analysis.Diagnostic{}
	// --- Read the flags --------------------------------------------------------
	a.once.Do(a.setup(pass.Analyzer.Flags))

	ins, _ := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// fmt.Println(pass.Pkg.Path())
	var (
		imports      = checker.Load(pass.Fset, ins)
		fileCheckers = make(map[string][]*checker.Checker)
		debugFn      = debugNoOp
		fileImports  []checker.Import
	)

	if a.withDebug {
		debugFn = debug(pass.Fset)
	}

	// --- Preorder Checker ------------------------------------------------------
	ins.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
		callExpr := n.(*ast.CallExpr)
		fileName := pass.Fset.Position(callExpr.Pos()).Filename

		if !a.withTests && strings.HasSuffix(fileName, "_test.go") {
			return
		}

		fileImports = imports.Lookup(fileName)
		if _, ok := fileCheckers[fileName]; !ok {
			fileCheckers[fileName] = a.filter(fileImports)
		}

		for i := range fileCheckers[fileName] {
			c := fileCheckers[fileName][i].With(pass, fileImports, debugFn)
			if violation := c.Check(callExpr); violation != nil {
				issues = append(issues, violation.Diagnostic(n.Pos(), n.End()))
				return
			}
		}
	})

	// --- Reporting issues ------------------------------------------------------
	for _, issue := range issues {
		pass.Report(*issue)
	}

	return nil, nil
}

func (a *analyzer) register(c *checker.Checker) {
	a.checkers[c.Package] = c
}

func (a *analyzer) filter(imports []checker.Import) []*checker.Checker {
	out := make([]*checker.Checker, 0, len(a.checkers))
	seen := make(map[string]bool, len(imports))

	var key string
	for i := range imports {
		key = imports[i].Pkg

		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = true

		if _, ok := a.checkers[key]; ok {
			out = append(out, a.checkers[key])
		}
	}

	return out
}

func (a *analyzer) setup(f flag.FlagSet) func() {
	return func() {
		a.withTests = f.Lookup("with-tests").Value.String() == "true"
		a.withDebug = f.Lookup("with-debug").Value.String() == "true"
	}
}

func flags() flag.FlagSet {
	set := flag.NewFlagSet("", flag.PanicOnError)
	set.Bool("with-tests", false, "do not skip tests in reports")
	set.Bool("with-debug", false, "debug linter run (development only)")
	return *set
}
