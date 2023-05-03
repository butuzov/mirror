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
	checkers []*checker.Checker        // Available checkers.
	imports  map[string]checker.Import // Map

	withTests bool
	withDebug bool

	lock sync.RWMutex
	once sync.Once
}

func NewAnalyzer() *analysis.Analyzer {
	flags := flags()

	a := analyzer{
		imports: map[string]checker.Import{},
		checkers: []*checker.Checker{
			newRegexpChecker(),
			newStringsChecker(),
			newBytesChecker(),
		},
	}

	return &analysis.Analyzer{
		Name: "mirror",
		Doc:  "looks for mirror patterns",
		Run:  a.run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
		Flags: flags,
	}
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	// --- Setup -----------------------------------------------------------------
	a.once.Do(a.setup(pass.Analyzer.Flags))

	ins, _ := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	imports := checker.LoadImports(pass.Fset, ins)
	issues := []*analysis.Diagnostic{}

	// --- Preorder Checker ------------------------------------------------------
	ins.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
		callExpr := n.(*ast.CallExpr)
		fileName := pass.Fset.Position(callExpr.Pos()).Filename

		// Reject tests if its not enabled.
		if !a.withTests && strings.HasSuffix(fileName, "_test.go") {
			return
		}

		fileImports := imports.LookupImports(fileName)

		// Run checkers against call expression.
		for _, check := range a.checkers {
			violation := check.With(pass.TypesInfo, fileImports).Check(callExpr)
			if violation != nil {
				issues = append(issues, violation.Diagnostic(n.Pos(), n.End()))
			}
		}
	})

	// --- Reporting issues ------------------------------------------------------
	for _, issue := range issues {
		pass.Report(*issue)
	}

	return nil, nil
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
