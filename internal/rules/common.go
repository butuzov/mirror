package rules

import (
	"go/ast"
	"go/types"
)

type Diagnostic struct {
	Message string
	Args    []int
}

func isBytesBufferType(tv types.TypeAndValue) bool {
	if !tv.IsValue() {
		return false
	}
	s := tv.Type.String()

	return s == "*bytes.Buffer" || s == "bytes.Buffer"
}

// return positions of args that matching to
func isBytesArrayCall(pos []int, args []ast.Expr) []int {
	var out []int

	for _, i := range pos {
		// todo(butuzov): check incoming variable type, if its variable

		// note(butuzov): checking call expression if its call expression.
		call, ok := args[i].(*ast.CallExpr)
		if !ok {
			continue
		}

		array, ok := call.Fun.(*ast.ArrayType)
		if !ok {
			continue
		}

		val, ok := array.Elt.(*ast.Ident)
		if !ok {
			continue
		}

		if val.Name != "byte" {
			continue
		}

		out = append(out, i)
	}

	return out
}

func isStringCall(pos []int, args []ast.Expr) []int {
	var out []int

	for _, i := range pos {
		// todo(butuzov): check incoming variable type, if its variable

		// note(butuzov): checking call expression if its call expression.
		call, ok := args[i].(*ast.CallExpr)
		if !ok {
			continue
		}

		funcName, ok := call.Fun.(*ast.Ident)
		if !ok {
			continue
		}

		if funcName.Name != "string" {
			continue
		}

		out = append(out, i)
	}

	return out
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
