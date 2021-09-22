package mirror

import (
	"go/ast"

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
		},
	}

	return &analysis.Analyzer{
		Name: "mirror",
		Doc:  "looks for mirror patterns",
		Run:  a.run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}
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

	for i := range issues {
		pass.Report(issues[i])
	}

	return nil, nil
}
