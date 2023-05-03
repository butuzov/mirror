package mirror

import "github.com/butuzov/mirror/internal/checker"

func newStringsChecker() *checker.Checker {
	c := checker.New("strings")
	c.Functions = StringFunctions
	c.Methods["strings.Builder"] = StringsBuilderMethods

	return c
}

var (
	StringFunctions = map[string]checker.Violation{
		"Compare": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.Compare",
			Args:           []int{0, 1},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `Compare`,
			},
			Generate: &checker.Generate{
				Pattern: `Compare($0,$1)`,
				Returns: 1,
			},
		},
		"Contains": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.Contains",
			Args:           []int{0, 1},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `Contains`,
			},
			Generate: &checker.Generate{
				Pattern: `Contains($0,$1)`,
				Returns: 1,
			},
		},
		"ContainsAny": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.ContainsAny",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `ContainsAny`,
			},
			Generate: &checker.Generate{
				Pattern: `ContainsAny($0,"foobar")`,
				Returns: 1,
			},
		},
		"ContainsRune": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.ContainsRune",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `ContainsRune`,
			},
			Generate: &checker.Generate{
				Pattern: `ContainsRune($0,'ф')`,
				Returns: 1,
			},
		},
		"Count": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.Count",
			Args:           []int{0, 1},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `Count`,
			},
			Generate: &checker.Generate{
				Pattern: `Count($0, $1)`,
				Returns: 1,
			},
		},
		"EqualFold": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.EqualFold",
			Args:           []int{0, 1},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `EqualFold`,
			},
			Generate: &checker.Generate{
				Pattern: `EqualFold($0,$1)`,
				Returns: 1,
			},
		},
		"HasPrefix": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.HasPrefix",
			Args:           []int{0, 1},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `HasPrefix`,
			},
			Generate: &checker.Generate{
				Pattern: `HasPrefix($0,$1)`,
				Returns: 1,
			},
		},
		"HasSuffix": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.HasSuffix",
			Args:           []int{0, 1},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `HasSuffix`,
			},
			Generate: &checker.Generate{
				Pattern: `HasSuffix($0,$1)`,
				Returns: 1,
			},
		},
		"Index": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.Index",
			Args:           []int{0, 1},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `Index`,
			},
			Generate: &checker.Generate{
				Pattern: `Index($0,$1)`,
				Returns: 1,
			},
		},
		"IndexAny": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.IndexAny",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `IndexAny`,
			},
			Generate: &checker.Generate{
				Pattern: `IndexAny($0, "f")`,
				Returns: 1,
			},
		},
		"IndexByte": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.IndexByte",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `IndexByte`,
			},
			Generate: &checker.Generate{
				Pattern: `IndexByte($0, byte('f'))`,
				Returns: 1,
			},
		},
		"IndexFunc": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.IndexFunc",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `IndexFunc`,
			},
			Generate: &checker.Generate{
				Pattern: `IndexFunc($0,func(r rune) bool { return true })`,
				Returns: 1,
			},
		},
		"IndexRune": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.IndexRune",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `IndexRune`,
			},
			Generate: &checker.Generate{
				Pattern: `IndexRune($0, rune('ф'))`,
				Returns: 1,
			},
		},
		"LastIndex": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.LastIndex",
			Args:           []int{0, 1},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `LastIndex`,
			},
			Generate: &checker.Generate{
				Pattern: `LastIndex($0,$1)`,
				Returns: 1,
			},
		},
		"LastIndexAny": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.LastIndexAny",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `LastIndexAny`,
			},
			Generate: &checker.Generate{
				Pattern: `LastIndexAny($0,"f")`,
				Returns: 1,
			},
		},
		"LastIndexByte": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.LastIndexByte",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `LastIndexByte`,
			},
			Generate: &checker.Generate{
				Pattern: `LastIndexByte($0, byte('f'))`,
				Returns: 1,
			},
		},
		"LastIndexFunc": {
			Type:           checker.Function,
			Message:        "avoid allocations with bytes.LastIndexAny",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  `bytes`,
				Function: `LastIndexAny`,
			},
			Generate: &checker.Generate{
				Pattern: `LastIndexFunc($0, func(r rune) bool { return true })`,
				Returns: 1,
			},
		},
	}

	StringsBuilderMethods = map[string]checker.Violation{
		"Write": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*strings.Builder).WriteString",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Method: `WriteString`,
			},
			Generate: &checker.Generate{
				PreCondition: `builder := strings.Builder{}`,
				Pattern:      `Write($0)`,
				Returns:      2,
			},
		},
		"WriteString": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*strings.Builder).Write",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Method: `Write`,
			},
			Generate: &checker.Generate{
				PreCondition: `builder := strings.Builder{}`,
				Pattern:      `WriteString($0)`,
				Returns:      2,
			},
		},
	}
)
