package mirror

import "github.com/butuzov/mirror/internal/checker"

func newUTF8Checker() *checker.Checker {
	c := checker.New("unicode/utf8")
	c.Functions = UTF8Functions

	return c
}

var UTF8Functions = map[string]checker.Violation{
	"Valid": {
		Type:           checker.Function,
		Message:        "avoid allocations with utf8.ValidString",
		Args:           []int{0},
		StringTargeted: false,
		Alternative: checker.Alternative{
			Package:  "unicode/utf8",
			Function: "ValidString",
		},
		Generate: &checker.Generate{
			Pattern: `Valid($0)`,
			Returns: 1,
		},
	},
	"ValidString": {
		Type:           checker.Function,
		Message:        "avoid allocations with utf8.Valid",
		Args:           []int{0},
		StringTargeted: true,
		Alternative: checker.Alternative{
			Package:  "unicode/utf8",
			Function: "Valid",
		},
		Generate: &checker.Generate{
			Pattern: `ValidString($0)`,
			Returns: 1,
		},
	},
}
