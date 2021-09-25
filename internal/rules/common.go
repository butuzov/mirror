package rules

import (
	"go/ast"
	"go/types"
)

type Diagnostic struct {
	Message       string
	Args          []int
	TargetStrings bool // Is this function/method expected to work with strings?

	GenCondition string // Precondition, if we working with struct (regexpRegexp or strings.Builder)
	GenPattern   string // Generated Call
	GenReturns   int    // Placeholder for generated test return results
}

func isBytesBufferType(tv types.TypeAndValue) bool {
	if !tv.IsValue() {
		return false
	}
	s := tv.Type.String()

	return s == "*bytes.Buffer" || s == "bytes.Buffer"
}

// Check will try to find which arguments can be replaced.
func check(pos []int, args []ast.Expr, isString bool) (matched []int) {
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
