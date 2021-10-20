package checker

import "go/ast"

type Converter struct {
	n *ast.CallExpr
	r ReturnType
}

func (c *Converter) Return() ReturnType {
	if c == nil {
		return typeUnknown
	}
	return c.r
}

func (c *Converter) Expr() *ast.CallExpr {
	if c == nil {
		return nil
	}
	return c.n
}

func (c *Converter) Valid() bool {
	return c != nil && (c.r == typeString || c.r == typeByteSlice)
}

// isConverter desides is this expression is conversion of []byte() or string()
func isConverter(n ast.Expr) (*Converter, bool) {
	call, ok := n.(*ast.CallExpr)
	if !ok {
		return &Converter{n: call, r: typeUnknown}, false
	}

	switch v := call.Fun.(type) {
	case *ast.ArrayType:
		if consistsOfBytes(v) {
			return &Converter{n: call, r: typeByteSlice}, true
		}

	case *ast.Ident:
		if indentIsString(v) {
			return &Converter{n: call, r: typeString}, true
		}
	}

	return &Converter{n: call, r: typeUnknown}, false
}
