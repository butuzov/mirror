package checker

import (
	"go/ast"
	"go/token"
	"go/types"
)

// Type checking of arguments we getting from string and []byte converters.
// e only care if undertype is string or []byte ignoring the rest.
// issue(butuzov): typesystem sometimes can't recognize strings, find out why.
// todo(butuzov): channel reading
// todo(butuzov): pointer dereference
func (rc *Checker) Type(n ast.Expr) ReturnType {
	switch node := n.(type) {

	case *ast.IndexExpr:
		return rc.handleIndexExpr(node)

	case *ast.BasicLit:
		return rc.handleBasicLiteral(node)

	case *ast.CompositeLit:
		return rc.handleCompositeLiteral(node)

	case *ast.CallExpr:
		return rc.handleCallExpression(node)

	case *ast.BinaryExpr:
		return typeString

	case *ast.SelectorExpr:
		return rc.handleSelectorExpression(node)

	case *ast.Ident:
		return rc.handleIdent(node)
	}

	rc.Debug(n)

	return typeUnknown
}

func (rc *Checker) handleCompositeLiteral(node *ast.CompositeLit) ReturnType {
	if isBytes(rc.types[node.Type].Type) {
		return typeByteSlice
	}

	rc.Debug(node)

	return typeUnknown
}

// handleCallExpression allow to find call expressions return type from its
// signature.
func (rc *Checker) handleCallExpression(node *ast.CallExpr) ReturnType {
	tv := rc.types[node.Fun]
	s, ok := tv.Type.(*types.Signature)
	if !ok {
		return typeUnknown
	}

	v := CallReturnsType(s)
	switch v {
	case typeByteSlice:
		return typeByteSlice
	case typeString:
		return typeString
	}

	rc.Debug(node)

	return typeUnknown
}

// handleSelectorExpression works for  selector expressions (word.Word)
//
// 		github.com/iser/project.Type
// 		some{Val string} -> some.Val
//
func (rc *Checker) handleSelectorExpression(node *ast.SelectorExpr) ReturnType {
	tv := rc.types[node]
	switch {
	case isVarString(tv.Type):
		return typeString
	case isVarBytes(tv.Type):
		return typeByteSlice
	case isNamedTypeOfBytes(tv.Type):
		return typeByteSlice
	case isNamedTypeOfString(tv.Type):
		return typeString
	case isRune(tv.Type):
		return typeString
	}

	rc.Debug(node)

	return typeUnknown
}

// handleBasicLiteral works for string literals we finding in out code,
// or other literals.
//
// 		string("string")
// 		[]byte("string")
//
func (rc *Checker) handleBasicLiteral(node *ast.BasicLit) ReturnType {
	if node.Kind == token.STRING {
		return typeString
	}

	tv := rc.types[node]
	if isVarString(tv.Type) {
		return typeString
	}

	rc.Debug(node)

	return typeUnknown
}

// handleIndexExpr allows to handle ast.Expressions that fall into
// category of index calls:
//
// 		map[string]string -> v["give_me_string"]
// 		[]string -> v[0]
// 		[2]string -> v[0]
//
func (rc *Checker) handleIndexExpr(node *ast.IndexExpr) ReturnType {
	tv := rc.types[node.X]

	switch {
	case isCollectionOfStrings(tv.Type):
		return typeString
	case isCollectionOfBytes(tv.Type):
		return typeByteSlice
	}

	rc.Debug(node)

	return typeUnknown
}

// handleIdent allow to ast.Indent to be checked, not reliable
// types isn't always can give correct information.
func (rc *Checker) handleIdent(node *ast.Ident) ReturnType {
	tv := rc.types[node]
	switch {
	case isVarString(tv.Type):
		return typeString
	case isVarBytes(tv.Type):
		return typeByteSlice
	case isNamedTypeOfBytes(tv.Type):
		return typeByteSlice
	case isNamedTypeOfString(tv.Type):
		return typeString
	case isRune(tv.Type):
		return typeString
	}

	rc.Debug(node)

	return typeUnknown
}
