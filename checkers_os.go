package mirror

import "github.com/butuzov/mirror/internal/checker"

func newOsChecker() *checker.Checker {
	c := checker.New("os")
	c.Methods["os.File"] = OsFileMethods

	return c
}

var OsFileMethods = map[string]checker.Violation{
	"Write": {
		Type:           checker.Method,
		Message:        "avoid allocations with (*os.File).WriteString",
		Args:           []int{0},
		StringTargeted: false,
		Alternative: checker.Alternative{
			Method: "WriteString",
		},
		Generate: &checker.Generate{
			PreCondition: `f := &os.File{}`,
			Pattern:      `Write($0)`,
			Returns:      2,
		},
	},
	"WriteString": {
		Type:           checker.Method,
		Message:        "avoid allocations with (*os.File).Write",
		Args:           []int{0},
		StringTargeted: true,
		Alternative: checker.Alternative{
			Method: "Write",
		},
		Generate: &checker.Generate{
			PreCondition: `f := &os.File{}`,
			Pattern:      `WriteString($0)`,
			Returns:      2,
		},
	},
}
