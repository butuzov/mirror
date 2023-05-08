package mirror

import "github.com/butuzov/mirror/internal/checker"

func newBufioChecker() *checker.Checker {
	c := checker.New("bufio")
	c.Methods["bufio.Writer"] = BufioMethods

	return c
}

var BufioMethods = map[string]checker.Violation{
	"Write": {
		Type:           checker.Method,
		Message:        "avoid allocations with (*bufio.Writer).WriteString",
		Args:           []int{0},
		StringTargeted: false,
		Alternative: checker.Alternative{
			Method: "WriteString",
		},
		Generate: &checker.Generate{
			PreCondition: `b := bufio.Writer{}`,
			Pattern:      `Write($0)`,
			Returns:      2,
		},
	},
	"WriteString": {
		Type:           checker.Method,
		Message:        "avoid allocations with (*bufio.Writer).Write",
		Args:           []int{0},
		StringTargeted: true,
		Alternative: checker.Alternative{
			Method: "Write",
		},
		Generate: &checker.Generate{
			PreCondition: `b := bufio.Writer{}`,
			Pattern:      `WriteString($0)`,
			Returns:      2,
		},
	},
}
