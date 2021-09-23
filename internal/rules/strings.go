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
		// "Compare": {
		// 	Message: "this call can be optimized with bytes.Compare",
		// 	Args:    []int{0, 1},
		// },
		// "Contains": {
		// 	Message: "this call can be optimized with bytes.Contains",
		// 	Args:    []int{0, 1},
		// },
		"ContainsAny": {
			Message: "this call can be optimized with bytes.ContainsAny",
			Args:    []int{0},
		},
		"ContainsRune": {
			Message: "this call can be optimized with bytes.ContainsRune",
			Args:    []int{0},
		},
		"Count": {
			Message: "this call can be optimized with bytes.Count",
			Args:    []int{0, 1},
		},
		"EqualFold": {
			Message: "this call can be optimized with bytes.EqualFold",
			Args:    []int{0, 1},
		},
		"HasPrefix": {
			Message: "this call can be optimized with bytes.HasPrefix",
			Args:    []int{0, 1},
		},
		"HasSuffix": {
			Message: "this call can be optimized with bytes.HasSuffix",
			Args:    []int{0, 1},
		},
		"Index": {
			Message: "this call can be optimized with bytes.Index",
			Args:    []int{0, 1},
		},
		"IndexAny": {
			Message: "this call can be optimized with bytes.IndexAny",
			Args:    []int{0},
		},
		"IndexByte": {
			Message: "this call can be optimized with bytes.IndexByte",
			Args:    []int{0},
		},
		"IndexFunc": {
			Message: "this call can be optimized with bytes.IndexFunc",
			Args:    []int{0},
		},
		"IndexRune": {
			Message: "this call can be optimized with bytes.IndexRune",
			Args:    []int{0},
		},
		"LastIndex": {
			Message: "this call can be optimized with bytes.LastIndex",
			Args:    []int{0, 1},
		},
		"LastIndexAny": {
			Message: "this call can be optimized with bytes.LastIndexAny",
			Args:    []int{0},
		},
		"Runes": {
			Message: "this call can be optimized with bytes.Runes",
			Args:    []int{0},
		},
		"LastIndexByte": {
			Message: "this call can be optimized with bytes.LastIndexByte",
			Args:    []int{0},
		},
		"LastIndexFunc": {
			Message: "this call can be optimized with bytes.LastIndexAny",
			Args:    []int{0},
		},
	}

	// note(butuzov): adding confidiance feature (flag and field) will allow to check other methods.
	stringsBuilderMethods = map[string]Diagnostic{
		"Write": {
			Message: "this call can be optimized with WriteString method",
			Args:    []int{0},
		},
		"WriteString": {
			Message: "this call can be optimized with Write method",
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
			res := isStringCall(d.Args, ce.Args)
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
			res := checkRegExp(d.Args, ce.Args, strings.Contains(v.Sel.Name, `String`))
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
			res := isStringCall(d.Args, ce.Args)
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
