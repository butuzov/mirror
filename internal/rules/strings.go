package rules

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/butuzov/mirror/internal/data"
	"golang.org/x/tools/go/analysis"
)

// todo(butuzov): detection of something that string(regexp)

type StringsCheckers struct{}

var (
	stringFunctions = map[string]Diagnostic{
		"Compare": {
			Message: "you should be using bytes.Compare",
			Args:    []int{0, 1},
		},
		"Contains": {
			Message: "you should be using bytes.Contains",
			Args:    []int{0, 1},
		},
	}

	// note(butuzov): adding confidiance feature (flag and field) will allow to check other methods.
	stringsBuilderMethods = map[string]Diagnostic{
		"Write": {
			Message: "you should be using WriteString method",
			Args:    []int{0},
		},
		"WriteString": {
			Message: "you should be using Write method",
			Args:    []int{0},
		}, //
	}
)

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
		if d, ok := stringFunctions[v.Sel.Name]; ok && ld.HasImport(`strings`, x.Name) {
			// proceed with check
			res := checkStringsExp(d.Args, ce.Args, strings.Contains(v.Sel.Name, `String`))
			if len(res) != len(d.Args) {
				return nil
			}

			// TODO: SuggestedFixes
			// 1) add or remove string
			// 2) remove parameter wrapping
			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}

		}

		// method of the regexp.Regexp
		if d, ok := stringsBuilderMethods[v.Sel.Name]; ok && isStringsBuilderType(ld.Types[v.X]) {
			// proceed with check
			res := checkStringsExp(d.Args, ce.Args, strings.Contains(v.Sel.Name, `String`))
			if len(res) != len(d.Args) {
				return nil
			}

			// TODO: SuggestedFixes
			// 1) add or remove string
			// 2) remove parameter wrapping
			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}

	case *ast.Ident:
		// function (v.Sel.Name) Match(String)? on dot imported package `regexp`
		if d, ok := stringFunctions[v.Name]; ok && ld.HasImport(`strings`, `.`) {
			// proceed with check
			res := checkStringsExp(d.Args, ce.Args, strings.Contains(v.Name, `String`))
			if len(res) != len(d.Args) {
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

// Check will try to find which arguments can be replaced.
func checkStringsExp(pos []int, args []ast.Expr, isString bool) (matched []int) {
	for _, i := range pos {
		call, ok := args[i].(*ast.CallExpr)
		if !ok {
			continue
		}

		switch node := call.Fun.(type) {
		case *ast.Ident:
			//
			// Converting with string()
			//
			// todo(butuzov): edge cases
			// - is it edge case fmt.Sprintf("%s", )
			//
			if isString != (node.Name != "string") {
				matched = append(matched, i)
				continue
			}
		case *ast.ArrayType:
			//
			// Converting to the bytes slice with direct convertion []byte
			//
			// todo(butuzov): edge cases
			// - argument suports .Bytes() []byte
			// - argument is []bytes (unnecessary converstion)
			//
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

	return matched
}
