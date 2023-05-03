package mirror

import "github.com/butuzov/mirror/internal/checker"

func newBytesChecker() *checker.Checker {
	c := checker.New("bytes")
	c.Functions = BytesFunctions
	c.Methods["bytes.Buffer"] = BytesBufferMethods

	return c
}

var (
	BytesFunctions = map[string]checker.Violation{
		"NewBuffer": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.NewBufferString",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "bytes",
				Function: "NewBufferString",
			},
			Generate: &checker.Generate{
				Pattern: `NewBuffer($0)`,
				Returns: 1,
			},
		},
		"NewBufferString": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.NewBuffer",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  "bytes",
				Function: "NewBuffer",
			},
			Generate: &checker.Generate{
				Pattern: `NewBufferString($0)`,
				Returns: 1,
			},
		},
		"Compare": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.Compare",
			Args:           []int{0, 1},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "Compare",
			},
			Generate: &checker.Generate{
				Pattern: `Compare($0, $1)`,
				Returns: 1,
			},
		},
		"Contains": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.Contains",
			Args:           []int{0, 1},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "Contains",
			},
			Generate: &checker.Generate{
				Pattern: `Contains($0, $1)`,
				Returns: 1,
			},
		},
		"ContainsAny": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.ContainsAny",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "ContainsAny",
			},
			Generate: &checker.Generate{
				Pattern: `ContainsAny($0, "f")`,
				Returns: 1,
			},
		},
		"ContainsRune": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.ContainsRune",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "ContainsRune",
			},
			Generate: &checker.Generate{
				Pattern: `ContainsRune($0, rune('ф'))`,
				Returns: 1,
			},
		},
		"Count": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.Count",
			Args:           []int{0, 1},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "Count",
			},
			Generate: &checker.Generate{
				Pattern: `Count($0, $1)`,
				Returns: 1,
			},
		},
		"EqualFold": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.EqualFold",
			Args:           []int{0, 1},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "EqualFold",
			},
			Generate: &checker.Generate{
				Pattern: `EqualFold($0, $1)`,
				Returns: 1,
			},
		},

		"HasPrefix": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.HasPrefix",
			Args:           []int{0, 1},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "HasPrefix",
			},
			Generate: &checker.Generate{
				Pattern: `HasPrefix($0, $1)`,
				Returns: 1,
			},
		},
		"HasSuffix": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.HasSuffix",
			Args:           []int{0, 1},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "HasSuffix",
			},
			Generate: &checker.Generate{
				Pattern: `HasSuffix($0, $1)`,
				Returns: 1,
			},
		},
		"Index": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.Index",
			Args:           []int{0, 1},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "Index",
			},
			Generate: &checker.Generate{
				Pattern: `Index($0, $1)`,
				Returns: 1,
			},
		},
		"IndexAny": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.IndexAny",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "IndexAny",
			},
			Generate: &checker.Generate{
				Pattern: `IndexAny($0, "f")`,
				Returns: 1,
			},
		},
		"IndexByte": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.IndexByte",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "IndexByte",
			},
			Generate: &checker.Generate{
				Pattern: `IndexByte($0, 'f')`,
				Returns: 1,
			},
		},
		"IndexFunc": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.IndexFunc",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "IndexFunc",
			},
			Generate: &checker.Generate{
				Pattern: `IndexFunc($0, func(rune) bool {return true })`,
				Returns: 1,
			},
		},
		"IndexRune": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.IndexRune",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "IndexRune",
			},
			Generate: &checker.Generate{
				Pattern: `IndexRune($0, rune('ф'))`,
				Returns: 1,
			},
		},
		"LastIndex": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.LastIndex",
			Args:           []int{0, 1},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "LastIndex",
			},
			Generate: &checker.Generate{
				Pattern: `LastIndex($0, $1)`,
				Returns: 1,
			},
		},
		"LastIndexAny": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.LastIndexAny",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "LastIndexAny",
			},
			Generate: &checker.Generate{
				Pattern: `LastIndexAny($0, "ф")`,
				Returns: 1,
			},
		},

		"LastIndexByte": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.LastIndexByte",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "LastIndexByte",
			},
			Generate: &checker.Generate{
				Pattern: `LastIndexByte($0, 'f')`,
				Returns: 1,
			},
		},
		"LastIndexFunc": {
			Type:           checker.Function,
			Message:        "avoid allocations with strings.LastIndexAny",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "strings",
				Function: "LastIndexAny",
			},
			Generate: &checker.Generate{
				Pattern: `LastIndexFunc($0, func(rune) bool {return true })`,
				Returns: 1,
			},
		},
	}

	BytesBufferMethods = map[string]checker.Violation{
		"Write": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*bytees.Buffer).WriteString",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Method: "WriteString",
			},
			Generate: &checker.Generate{
				PreCondition: `bb := bytes.Buffer{}`,
				Pattern:      `Write($0)`,
				Returns:      2,
			},
		},
		"WriteString": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*bytees.Buffer).Write",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Method: "Write",
			},
			Generate: &checker.Generate{
				PreCondition: `bb := bytes.Buffer{}`,
				Pattern:      `WriteString($0)`,
				Returns:      2,
			},
		},
	}
)
