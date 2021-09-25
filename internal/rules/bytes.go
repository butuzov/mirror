package rules

import (
	"go/ast"

	"github.com/butuzov/mirror/internal/data"
	"golang.org/x/tools/go/analysis"
)

type BytesCheckers struct{}

func (re *BytesCheckers) Check(ce *ast.CallExpr, ld *data.Data) []analysis.Diagnostic {
	switch v := ce.Fun.(type) {
	case *ast.SelectorExpr:

		x, ok := v.X.(*ast.Ident)
		if !ok {
			return nil
		}

		if d, ok := BytesFunctions[v.Sel.Name]; ok && ld.HasImport(`bytes`, x.Name) {
			// fmt.Println(v.Sel.Name, ce.Args)
			if res := check(d.Args, ce.Args, d.TargetStrings); len(res) != len(d.Args) {
				return nil
			}
			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}

		// method of the regexp.Regexp
		if d, ok := BytesBufferMethods[v.Sel.Name]; ok && isBytesBufferType(ld.Types[v.X]) {
			// proceed with check
			res := check(d.Args, ce.Args, d.TargetStrings)
			if len(res) != len(d.Args) {
				return nil
			}

			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}

	case *ast.Ident:
		if d, ok := BytesFunctions[v.Name]; ok && ld.HasImport(`bytes`, `.`) {

			if res := check(d.Args, ce.Args, d.TargetStrings); len(res) != len(d.Args) {
				return nil
			}

			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}
	}

	return nil
}
