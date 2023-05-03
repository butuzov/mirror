package checker

type Type string

const (
	String Type = "string"
	Bytes  Type = "[]byte"
)

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
