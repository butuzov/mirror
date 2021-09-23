package rules

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/butuzov/mirror/internal/data"
	"golang.org/x/tools/go/analysis"
)

// This rules are regarding package regexp and allow to check for the
// next patterns:
//
// regexp.Match(".*", []byte("string")) -> regexp.MatchString(".*", "string")

// todo(butuzov): detection of something that string(regexp)

type RegexpCheckers struct{}

var (
	RegexpFunctions = map[string]Diagnostic{
		"Match": {
			Message: "this call can be optimized with regexp.MatchString",
			Args:    []int{1},
		},
		"MatchString": {
			Message: "this call can be optimized with regexp.Match",
			Args:    []int{1},
		},
	}

	// As you see we are not using all of the regexp method because
	// nes we missing return concrete types (bytes or strings)
	// which most probably was intentional.

	// note(butuzov): adding confidiance feature (flag and field) will allow to check other methods.
	RegexpMethods = map[string]Diagnostic{
		"Match": {
			Message: "this call can be optimized with (*regexp.Regexp).MatchString",
			Args:    []int{0},
		},
		"MatchString": {
			Message: "this call can be optimized with (*regexp.Regexp).Match",
			Args:    []int{0},
		},
		"FindAllIndex": {
			Message: "this call can be optimized with (*regexp.Regexp).FindAllStringIndex",
			Args:    []int{0},
		},
		"FindAllStringIndex": {
			Message: "this call can be optimized with (*regexp.Regexp).FindAllIndex",
			Args:    []int{0},
		},
		"FindAllSubmatchIndex": {
			Message: "this call can be optimized with (*regexp.Regexp).FindAllStringSubmatchIndex",
			Args:    []int{0},
		}, //
		"FindAllStringSubmatchIndex": {
			Message: "this call can be optimized with (*regexp.Regexp).FindAllSubmatchIndex",
			Args:    []int{0},
		},
		"FindIndex": {
			Message: "this call can be optimized with (*regexp.Regexp).FindStringIndex",
			Args:    []int{0},
		},
		"FindStringIndex": {
			Message: "this call can be optimized with (*regexp.Regexp).FindStringIndex",
			Args:    []int{0},
		},
		"FindSubmatchIndex": {
			Message: "this call can be optimized with (*regexp.Regexp).FindStringSubmatchIndex",
			Args:    []int{0},
		},
		"FindStringSubmatchIndex": {
			Message: "this call can be optimized with (*regexp.Regexp).FindSubmatchIndex",
			Args:    []int{0},
		},
	}
)

func (re *RegexpCheckers) Check(ce *ast.CallExpr, ld *data.Data) []analysis.Diagnostic {
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
		if d, ok := RegexpFunctions[v.Sel.Name]; ok && ld.HasImport(`regexp`, x.Name) {
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

		// method of the regexp.Regexp
		if d, ok := RegexpMethods[v.Sel.Name]; ok && isRegexpRegexpType(ld.Types[v.X]) {
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
		if d, ok := RegexpFunctions[v.Name]; ok && ld.HasImport(`regexp`, `.`) {
			// proceed with check
			res := checkRegExp(d.Args, ce.Args, strings.Contains(v.Name, `String`))
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

func isRegexpRegexpType(tv types.TypeAndValue) bool {
	if !tv.IsValue() {
		return false
	}
	s := tv.Type.String()

	return s == "*regexp.Regexp" || s == "regexp.Regexp"
}
