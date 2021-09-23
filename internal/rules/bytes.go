package rules

import (
	"go/ast"
	"strings"

	"github.com/butuzov/mirror/internal/data"
	"golang.org/x/tools/go/analysis"
)

type BytesCheckers struct{}

var (
	bytesFunctions = map[string]Diagnostic{
		"NewBuffer": {
			Message: "this call can be optimized with bytes.NewBufferString",
			Args:    []int{0},
		},
		"NewBufferString": {
			Message: "this call can be optimized with bytes.NewBuffer",
			Args:    []int{0},
		},
		"Compare": {
			Message: "this call can be optimized with strings.Compare",
			Args:    []int{0, 1},
		},
		"Contains": {
			Message: "this call can be optimized with strings.Contains",
			Args:    []int{0, 1},
		},
		"ContainsAny": {
			Message: "this call can be optimized with strings.ContainsAny",
			Args:    []int{0},
		},
		"ContainsRune": {
			Message: "this call can be optimized with strings.ContainsRune",
			Args:    []int{0},
		},
		"Count": {
			Message: "this call can be optimized with strings.Count",
			Args:    []int{0, 1},
		},
		"EqualFold": {
			Message: "this call can be optimized with strings.EqualFold",
			Args:    []int{0, 1},
		},
		"HasPrefix": {
			Message: "this call can be optimized with strings.HasPrefix",
			Args:    []int{0, 1},
		},
		"HasSuffix": {
			Message: "this call can be optimized with strings.HasSuffix",
			Args:    []int{0, 1},
		},
		"Index": {
			Message: "this call can be optimized with strings.Index",
			Args:    []int{0, 1},
		},
		"IndexAny": {
			Message: "this call can be optimized with strings.IndexAny",
			Args:    []int{0},
		},
		"IndexByte": {
			Message: "this call can be optimized with strings.IndexByte",
			Args:    []int{0},
		},
		"IndexFunc": {
			Message: "this call can be optimized with strings.IndexFunc",
			Args:    []int{0},
		},
		"IndexRune": {
			Message: "this call can be optimized with strings.IndexRune",
			Args:    []int{0},
		},
		"LastIndex": {
			Message: "this call can be optimized with strings.LastIndex",
			Args:    []int{0, 1},
		},
		"LastIndexAny": {
			Message: "this call can be optimized with strings.LastIndexAny",
			Args:    []int{0},
		},
		"Runes": {
			Message: "this call can be optimized with strings.Runes",
			Args:    []int{0},
		},
		"LastIndexByte": {
			Message: "this call can be optimized with strings.LastIndexByte",
			Args:    []int{0},
		},
		"LastIndexFunc": {
			Message: "this call can be optimized with strings.LastIndexAny",
			Args:    []int{0},
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
			// fmt.Println(v.Sel.Name, ce.Args)
			if res := isBytesArrayCall(d.Args, ce.Args); len(res) != len(d.Args) {
				return nil
			}
			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}

		// method of the regexp.Regexp
		if d, ok := bytesBufferMethods[v.Sel.Name]; ok && isBytesBufferType(ld.Types[v.X]) {
			// proceed with check
			res := checkBytesExp(d.Args, ce.Args, strings.Contains(v.Sel.Name, `String`))
			if len(res) != len(d.Args) {
				return nil
			}

			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}

	case *ast.Ident:
		if d, ok := bytesFunctions[v.Name]; ok && ld.HasImport(`bytes`, `.`) {
			// fmt.Println(v.Name)
			if res := isBytesArrayCall(d.Args, ce.Args); len(res) != len(d.Args) {
				return nil
			}

			return []analysis.Diagnostic{{Pos: ce.Pos(), Message: d.Message}}
		}
	}

	return nil
}

// Check will try to find which arguments can be replaced.
func checkBytesExp(pos []int, args []ast.Expr, isString bool) (matched []int) {
	for _, i := range pos {
		call, ok := args[i].(*ast.CallExpr)
		if !ok {
			continue
		}

		switch node := call.Fun.(type) {
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
