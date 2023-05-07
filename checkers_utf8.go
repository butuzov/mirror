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
	"FullRune": {
		Type:           checker.Function,
		Message:        "avoid allocations with utf8.FullRuneInString",
		Args:           []int{0},
		StringTargeted: false,
		Alternative: checker.Alternative{
			Package:  "unicode/utf8",
			Function: "FullRuneInString",
		},
		Generate: &checker.Generate{
			Pattern: `FullRune($0)`,
			Returns: 1,
		},
	},
	"FullRuneInString": {
		Type:           checker.Function,
		Message:        "avoid allocations with utf8.FullRune",
		Args:           []int{0},
		StringTargeted: true,
		Alternative: checker.Alternative{
			Package:  "unicode/utf8",
			Function: "FullRune",
		},
		Generate: &checker.Generate{
			Pattern: `FullRuneInString($0)`,
			Returns: 1,
		},
	},

	"RuneCount": {
		Type:           checker.Function,
		Message:        "avoid allocations with utf8.RuneCountInString",
		Args:           []int{0},
		StringTargeted: false,
		Alternative: checker.Alternative{
			Package:  "unicode/utf8",
			Function: "RuneCountInString",
		},
		Generate: &checker.Generate{
			Pattern: `RuneCount($0)`,
			Returns: 1,
		},
	},
	"RuneCountInString": {
		Type:           checker.Function,
		Message:        "avoid allocations with utf8.RuneCount",
		Args:           []int{0},
		StringTargeted: true,
		Alternative: checker.Alternative{
			Package:  "unicode/utf8",
			Function: "RuneCount",
		},
		Generate: &checker.Generate{
			Pattern: `RuneCountInString($0)`,
			Returns: 1,
		},
	},

	"DecodeLastRune": {
		Type:           checker.Function,
		Message:        "avoid allocations with utf8.DecodeLastRuneInString",
		Args:           []int{0},
		StringTargeted: false,
		Alternative: checker.Alternative{
			Package:  "unicode/utf8",
			Function: "DecodeLastRuneInString",
		},
		Generate: &checker.Generate{
			Pattern: `DecodeLastRune($0)`,
			Returns: 2,
		},
	},
	"DecodeLastRuneInString": {
		Type:           checker.Function,
		Message:        "avoid allocations with utf8.DecodeLastRune",
		Args:           []int{0},
		StringTargeted: true,
		Alternative: checker.Alternative{
			Package:  "unicode/utf8",
			Function: "DecodeLastRune",
		},
		Generate: &checker.Generate{
			Pattern: `DecodeLastRuneInString($0)`,
			Returns: 2,
		},
	},
	"DecodeRune": {
		Type:           checker.Function,
		Message:        "avoid allocations with utf8.DecodeRuneInString",
		Args:           []int{0},
		StringTargeted: false,
		Alternative: checker.Alternative{
			Package:  "unicode/utf8",
			Function: "DecodeRuneInString",
		},
		Generate: &checker.Generate{
			Pattern: `DecodeRune($0)`,
			Returns: 2,
		},
	},
	"DecodeRuneInString": {
		Type:           checker.Function,
		Message:        "avoid allocations with utf8.DecodeRune",
		Args:           []int{0},
		StringTargeted: true,
		Alternative: checker.Alternative{
			Package:  "unicode/utf8",
			Function: "DecodeRune",
		},
		Generate: &checker.Generate{
			Pattern: `DecodeRuneInString($0)`,
			Returns: 2,
		},
	},
}
