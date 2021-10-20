package checker

import "go/types"

type ReturnType int

const (
	typeUnknown ReturnType = iota
	typeByte
	typeString
	typeByteSlice
)

func (t ReturnType) String() string {
	switch t {
	case typeByte:
		return `byte`
	case typeString:
		return `string`
	case typeByteSlice:
		return `[]byte`
	}

	return "unknown"
}

func isCollectionOfStrings(t types.Type) bool {
	col, ok := t.(interface{ Elem() types.Type })
	if !ok {
		return false
	}

	kind, ok := col.Elem().Underlying().(*types.Basic)
	if !ok {
		return false
	}

	return isStringType(kind.Kind())
}

func isCollectionOfBytes(t types.Type) bool {
	col, ok := t.(interface{ Elem() types.Type })
	if !ok {
		return false
	}

	return isBytes(col.Elem())
}

func isBytes(t types.Type) bool {
	col, ok := t.(interface{ Elem() types.Type })
	if !ok {
		return false
	}

	b, ok := col.Elem().(*types.Basic)
	if !ok {
		return false
	}

	return isByteType(b.Kind())
}

func isVarString(t types.Type) bool {
	v, ok := t.(*types.Basic)
	if !ok {
		return false
	}

	return isStringType(v.Kind())
}

func isRune(t types.Type) bool {
	v, ok := t.(*types.Basic)
	if !ok {
		return false
	}

	return isRuneType(v.Kind())
}

func isVarBytes(t types.Type) bool {
	v, ok := t.(*types.Slice)
	if !ok {
		return false
	}

	b, ok := v.Elem().Underlying().(*types.Basic)
	if !ok {
		return false
	}

	return isByteType(b.Kind())
}

func isNamedTypeOfBytes(t types.Type) bool {
	tn, ok := t.(*types.Named)
	if !ok {
		return false
	}
	return isBytes(tn.Underlying())
}

func isNamedTypeOfString(t types.Type) bool {
	tn, ok := t.(*types.Named)
	if !ok {
		return false
	}
	return isVarString(tn.Underlying())
}

func CallReturnsType(s *types.Signature) ReturnType {
	if s.Results().Len() != 1 {
		// return more than 1 result
		return typeUnknown
	}

	t := s.Results().At(0).Type()
	switch {
	case isBytes(t):
		return typeByteSlice
	case isVarString(t.Underlying()):
		return typeString
	case isNamedTypeOfBytes(t):
		return typeByteSlice
	case isNamedTypeOfString(t):
		return typeString
	}

	return typeUnknown
}

func isStringType(k types.BasicKind) bool {
	return k == types.String || k == types.UntypedString
}

func isByteType(k types.BasicKind) bool {
	return k == types.Byte
}

func isRuneType(k types.BasicKind) bool {
	return k == types.Rune || k == types.UntypedRune
}
