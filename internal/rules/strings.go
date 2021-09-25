package rules

import (
	"go/ast"
	"go/types"

	"github.com/butuzov/mirror/internal/data"
	"golang.org/x/tools/go/analysis"
)

// todo(butuzov): detection of something that string(regexp)

type StringsCheckers struct{}

func (re *StringsCheckers) Check(ce *ast.CallExpr, ld *data.Data) []analysis.Diagnostic {
	switch v := ce.Fun.(type) {
	case *ast.SelectorExpr:

		// selector expression can be matched to package functions (regexp.Match) &
		// also tot he receiver methods, so first we need to know which of this
		// two cases we looking for.
		x, ok := v.X.(*ast.Ident)
		if !ok {
			return nil
		}

		// function (v.Sel.Name) Match(String)? on imported package `regexp` (x.Name)
		if d, ok := StringFunctions[v.Sel.Name]; ok && ld.HasImport(`strings`, x.Name) {
			// proceed with check
			if res := check(d.Args, ce.Args, d.TargetStrings); len(res) != len(d.Args) {
				return nil
			}

			// TODO: SuggestedFixes
			// 1) add or remove string
			// 2) remove parameter wrapping
			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}

		}

		// method of the regexp.Regexp
		if d, ok := StringsBuilderMethods[v.Sel.Name]; ok && isStringsBuilderType(ld.Types[v.X]) {
			// proceed with check
			if res := check(d.Args, ce.Args, d.TargetStrings); len(res) != len(d.Args) {
				return nil
			}

			// TODO: SuggestedFixes
			// 1) add or remove string
			// 2) remove parameter wrapping
			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}

	case *ast.Ident:
		// function (v.Sel.Name) Match(String)? on dot imported package `regexp`
		if d, ok := StringFunctions[v.Name]; ok && ld.HasImport(`strings`, `.`) {
			// proceed with check
			if res := check(d.Args, ce.Args, d.TargetStrings); len(res) != len(d.Args) {
				return nil
			}

			// TODO: SuggestedFixes
			// 1) add or remove string
			// 2) remove parameter wrapping
			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}
	}

	return nil
}

func isStringsBuilderType(tv types.TypeAndValue) bool {
	if !tv.IsValue() {
		return false
	}
	s := tv.Type.String()

	return s == "*strings.Builder" || s == "strings.Builder"
}
