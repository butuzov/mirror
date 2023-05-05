package mirror

import "github.com/butuzov/mirror/internal/checker"

func newMaphashChecker() *checker.Checker {
	c := checker.New("hash/maphash")
	c.Methods["hash/maphash.Hash"] = MaphashMethods

	return c
}

var MaphashMethods = map[string]checker.Violation{
	"Write": {
		Type:           checker.Method,
		Message:        "avoid allocations with (*maphash.Hash).WriteString",
		Args:           []int{0},
		StringTargeted: false,
		Alternative: checker.Alternative{
			Method: "WriteString",
		},
		Generate: &checker.Generate{
			PreCondition: `h := maphash.Hash{}`,
			Pattern:      `Write($0)`,
			Returns:      2,
		},
	},
	"WriteString": {
		Type:           checker.Method,
		Message:        "avoid allocations with (*maphash.Hash).Write",
		Args:           []int{0},
		StringTargeted: true,
		Alternative: checker.Alternative{
			Method: "Write",
		},
		Generate: &checker.Generate{
			PreCondition: `h := maphash.Hash{}`,
			Pattern:      `WriteString($0)`,
			Returns:      2,
		},
	},
}
