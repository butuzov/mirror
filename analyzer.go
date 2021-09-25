package mirror

import (
	"flag"
	"go/ast"
	"strings"

	"github.com/butuzov/mirror/internal/data"
	"github.com/butuzov/mirror/internal/rules"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type checker interface {
	Check(*ast.CallExpr, *data.Data) []analysis.Diagnostic
}

type analyzer struct {
	imports  map[string][]string // a map of aliases of the packages
	checkers []checker
}

func NewAnalyzer() *analysis.Analyzer {
	a := analyzer{
		imports: map[string][]string{},
		checkers: []checker{
			&rules.RegexpCheckers{},
			&rules.StringsCheckers{},
			&rules.BytesCheckers{},
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
	return *set
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	ins, _ := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	ld := data.Data{}
	ld.MapExistingImports(ins)
	ld.MapTypesInfo(pass.TypesInfo.Types)

	issues := []analysis.Diagnostic{}

	ins.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
		node, _ := n.(*ast.CallExpr)
		for _, v := range a.checkers {
			issues = append(issues,
				v.Check(node, &ld)...)
		}
	})

	withTests := pass.Analyzer.Flags.Lookup("with-tests").Value.String() == "true"

	for i := range issues {
		if !withTests && isTest(pass.Fset.PositionFor(issues[i].Pos, true).Filename) {
			continue
		}
		pass.Report(issues[i])
	}

	return nil, nil
}

func isTest(fileName string) bool {
	return strings.HasSuffix(fileName, "_test.go")
}
