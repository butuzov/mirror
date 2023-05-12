package mirror

import "github.com/butuzov/mirror/internal/checker"

var BufioMethods = []checker.Violation{
	{ // (*bufio.Writer).Write
		Targets:   checker.Bytes,
		Type:      checker.Method,
		Package:   "bufio",
		Struct:    "Writer",
		Caller:    "Write",
		Args:      []int{0},
		AltCaller: "WriteString",

		Generate: &checker.Generate{
			PreCondition: `b := bufio.Writer{}`,
			Pattern:      `Write($0)`,
			Returns:      2,
		},
	},
	{ // (*bufio.Writer).WriteString
		Type:      checker.Method,
		Targets:   checker.Strings,
		Package:   "bufio",
		Struct:    "Writer",
		Caller:    "WriteString",
		Args:      []int{0},
		AltCaller: "Write",

		Generate: &checker.Generate{
			PreCondition: `b := bufio.Writer{}`,
			Pattern:      `WriteString($0)`,
			Returns:      2,
		},
	},
}
