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

type Diagnostic struct {
	Message string
	Args    []int
}

var (
	regexpFunctions = map[string]Diagnostic{
		"Match": {
			Message: "you should be using regexp.MatchString",
			Args:    []int{1},
		},
		"MatchString": {
			Message: "you should be using regexp.Match",
			Args:    []int{1},
		},
	}

	// As you see we are not using all of the regexp method because
	// nes we missing return concrete types (bytes or strings)
	// which most probably was intentional.

	// note(butuzov): adding confidiance feature (flag and field) will allow to check other methods.
	regexpMethods = map[string]Diagnostic{
		"Match": {
			Message: "you should be using MatchString method",
			Args:    []int{0},
		},
		"MatchString": {
			Message: "you should be using Match method",
			Args:    []int{0},
		},
		"FindAllIndex": {
			Message: "you should be using FindAllStringIndex method",
			Args:    []int{0},
		},
		"FindAllStringIndex": {
			Message: "you should be using FindAllIndex method",
			Args:    []int{0},
		},
		"FindAllSubmatchIndex": {
			Message: "you should be using FindAllStringSubmatchIndex method",
			Args:    []int{0},
		}, //
		"FindAllStringSubmatchIndex": {
			Message: "you should be using FindAllSubmatchIndex method",
			Args:    []int{0},
		},
		"FindIndex": {
			Message: "you should be using FindStringIndex method",
			Args:    []int{0},
		},
		"FindStringIndex": {
			Message: "you should be using FindStringIndex method",
			Args:    []int{0},
		},
		"FindSubmatchIndex": {
			Message: "you should be using FindStringSubmatchIndex method",
			Args:    []int{0},
		},
		"FindStringSubmatchIndex": {
			Message: "you should be using FindSubmatchIndex method",
			Args:    []int{0},
		},
	}
)

type typeLookupMap = *map[ast.Expr]types.TypeAndValue

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
		if d, ok := regexpFunctions[v.Sel.Name]; ok && ld.HasImport(`regexp`, x.Name) {
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
		if d, ok := regexpMethods[v.Sel.Name]; ok && isRegexpRegexpType(ld.Types[v.X]) {
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
		if d, ok := regexpFunctions[v.Name]; ok && ld.HasImport(`regexp`, `.`) {
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

// Check will try to find which arguments can be replaced.
func checkRegExp(pos []int, args []ast.Expr, isString bool) (matched []int) {
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
