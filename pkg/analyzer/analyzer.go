package analyzer

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const targetType = `*regexp.Regexp`

type analyzer struct {
	found    []analysis.Diagnostic
	packages []string
}

func NewAnalyzer() *analysis.Analyzer {
	a := analyzer{}

	return &analysis.Analyzer{
		Name:     "mirror",
		Doc:      "looks for mirror patterns",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	ins, _ := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// In order to minimize false positives we need to check if regexp
	// package was imported, and if so, get its names.

	ins.Preorder([]ast.Node{(*ast.ImportSpec)(nil)}, func(node ast.Node) {
		is, _ := node.(*ast.ImportSpec)
		if is.Path.Value == `"regexp"` {
			if is.Name == nil {
				a.packages = append(a.packages, "regexp")
				return
			}
			a.packages = append(a.packages, is.Name.Name)
		}
	})

	ins.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
		node, _ := n.(*ast.CallExpr)
		switch v := node.Fun.(type) {
		case *ast.SelectorExpr:

			x, ok := v.X.(*ast.Ident)
			if !ok {
				return
			}

			// function Match(String)? on imported package `regexp`
			if isFunction(v.Sel.Name) && a.hasImport(x.Name) {
				function := v.Sel.Name

				found := proceed(regexpFunctions[function], node.Args, isStringFunction(function))
				if len(found) == 0 {
					return
				}

				a.found = append(a.found, analysis.Diagnostic{
					Pos:     n.Pos(),
					Message: fmt.Sprintf("%s can be (insert alternative)", function),
				})
			}

			// method of `regexp.Regexp`

			if isMethod(v.Sel.Name) && isRegExp(pass.TypesInfo.Types[v.X]) {
				method := v.Sel.Name
				found := proceed(regexpMethods[method], node.Args, isStringFunction(method))
				if len(found) == 0 {
					return
				}

				a.found = append(a.found, analysis.Diagnostic{
					Pos:     n.Pos(),
					Message: fmt.Sprintf("%s can be (insert alternative)", method),
				})

				return
			}

		case *ast.Ident:

			// function Match(String)? on dot imported package `regexp`
			if isFunction(v.Name) && a.hasImport(dot) {
				function := v.Name

				found := proceed(regexpFunctions[function], node.Args, isStringFunction(function))
				if len(found) == 0 {
					return
				}

				a.found = append(a.found, analysis.Diagnostic{
					Pos:     n.Pos(),
					Message: fmt.Sprintf("%s can be (insert alternative)", function),
				})

				return
			}
		}
	})

	// 02. Printing reports.
	for i := range a.found {
		pass.Report(a.found[i])
	}

	return nil, nil
}

// proceedFunctionCall will try to find out if elements of
func proceed(pos []int, args []ast.Expr, isString bool) (matched []int) {
	for _, i := range pos {
		expr, ok := args[i].(*ast.CallExpr)
		if !ok {
			continue
		}

		switch node := expr.Fun.(type) {
		case *ast.Ident:
			// Possible scenarios
			// + string() can convert not only []byte() <- this pattern we trying ot cover
			// - string() can convert something else []uint8 <- same as byte
			// - string() can convert int
			// - string() can convert retult of other function call (so we need to an wind stack)
			if isString != (node.Name != "string") {
				matched = append(matched, i)
				continue
			}

		case *ast.ArrayType:
			val, ok := node.Elt.(*ast.Ident)
			if !ok {
				continue
			}

			if !isString != (val.Name != "byte") {
				matched = append(matched, i)
				continue
			}
		}

	}

	return
}

func isStringFunction(name string) bool {
	return strings.Contains(name, "String")
}

const dot string = "."

// hasImport is checking if slice of names for regexp imports has name in it.
func (a *analyzer) hasImport(name string) bool {
	for i := range a.packages {
		if a.packages[i] == name {
			return true
		}
	}
	return false
}

func isFunction(name string) bool {
	_, ok := regexpFunctions[name]
	return ok
}

func isMethod(name string) bool {
	_, ok := regexpMethods[name]
	return ok
}

func isRegExp(tv types.TypeAndValue) bool {
	if !tv.IsValue() {
		return false
	}

	if tv.Type.String() != targetType {
		return false
	}

	return true
}
