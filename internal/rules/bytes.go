package rules

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/butuzov/mirror/internal/data"
	"golang.org/x/tools/go/analysis"
)

type BytesCheckers struct{}

var (
	bytesFunctions = map[string]Diagnostic{
		"NewBuffer": {
			Message: "you should be using bytes.NewBufferString",
			Args:    []int{0},
		},
		"NewBufferString": {
			Message: "you should be using bytes.NewBuffer",
			Args:    []int{0},
		},
		"Compare": {
			Message: "you should be using strings.Compare",
			Args:    []int{0, 1},
		},
		"Contains": {
			Message: "you should be using strings.Contains",
			Args:    []int{0, 1},
		},
	}

	// note(butuzov): adding confidiance feature (flag and field) will allow to check other methods.
	bytesBufferMethods = map[string]Diagnostic{
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

func (re *BytesCheckers) Check(ce *ast.CallExpr, ld *data.Data) []analysis.Diagnostic {
	switch v := ce.Fun.(type) {
	case *ast.SelectorExpr:

		x, ok := v.X.(*ast.Ident)
		if !ok {
			return nil
		}

		if d, ok := bytesFunctions[v.Sel.Name]; ok && ld.HasImport(`bytes`, x.Name) {
			res := checkBytesExp(d.Args, ce.Args, strings.Contains(v.Sel.Name, `String`))
			if len(res) != len(d.Args) {
				return nil
			}

			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}

		// method of the regexp.Regexp
		if d, ok := stringsBuilderMethods[v.Sel.Name]; ok && isBytesBufferType(ld.Types[v.X]) {
			// proceed with check
			res := checkBytesExp(d.Args, ce.Args, strings.Contains(v.Sel.Name, `String`))
			if len(res) != len(d.Args) {
				return nil
			}

			// TODO: SuggestedFixes
			// 1) add or remove string
			// 2) remove parameter wrapping
			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}

	case *ast.Ident:
		if d, ok := bytesFunctions[v.Name]; ok && ld.HasImport(`bytes`, `.`) {
			res := checkBytesExp(d.Args, ce.Args, strings.Contains(v.Name, `String`))
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

func isBytesBufferType(tv types.TypeAndValue) bool {
	if !tv.IsValue() {
		return false
	}
	s := tv.Type.String()

	return s == "*bytes.Buffer" || s == "bytes.Buffer"
}

// Check will try to find which arguments can be replaced.
func checkBytesExp(pos []int, args []ast.Expr, isString bool) (matched []int) {
	for _, i := range pos {
		call, ok := args[i].(*ast.CallExpr)
		if !ok {
			continue
		}

		switch node := call.Fun.(type) {
		case *ast.Ident:

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

	return matched
}
