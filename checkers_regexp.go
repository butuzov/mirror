package mirror

import "github.com/butuzov/mirror/internal/checker"

func newRegexpChecker() *checker.Checker {
	c := checker.New("regexp")
	c.Functions = RegexpFunctions
	c.Methods["regexp.Regexp"] = RegexpRegexpMethods

	return c
}

var (
	RegexpFunctions = map[string]checker.Violation{
		"Match": {
			Type:           checker.Function,
			Message:        "avoid allocations with regexp.MatchString",
			Args:           []int{1},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Package:  "regexp",
				Function: "MatchString",
			},
			Generate: &checker.Generate{
				Pattern: `Match("foo", $0)`,
				Returns: 2,
			},
		},
		"MatchString": {
			Type:           checker.Function,
			Message:        "avoid allocations with regexp.Match",
			Args:           []int{1},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Package:  "regexp",
				Function: "Match",
			},
			Generate: &checker.Generate{
				Pattern: `MatchString("foo", $0)`,
				Returns: 2,
			},
		},
	}

	// As you see we are not using all of the regexp method because
	// nes we missing return concrete types (bytes or strings)
	// which most probably was intentional.
	RegexpRegexpMethods = map[string]checker.Violation{
		"Match": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*regexp.Regexp).MatchString",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Method: "MatchString",
			},
			Generate: &checker.Generate{
				PreCondition: `re := regexp.MustCompile(".*")`,
				Pattern:      `Match($0)`,
				Returns:      1,
			},
		},
		"MatchString": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*regexp.Regexp).Match",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Method: "Match",
			},
			Generate: &checker.Generate{
				PreCondition: `re := regexp.MustCompile(".*")`,
				Pattern:      `MatchString($0)`,
				Returns:      1,
			},
		},
		"FindAllIndex": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*regexp.Regexp).FindAllStringIndex",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Method: "FindAllStringIndex",
			},
			Generate: &checker.Generate{
				PreCondition: `re := regexp.MustCompile(".*")`,
				Pattern:      `FindAllIndex($0, 1)`,
				Returns:      1,
			},
		},
		"FindAllStringIndex": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*regexp.Regexp).FindAllIndex",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Method: "FindAllIndex",
			},
			Generate: &checker.Generate{
				PreCondition: `re := regexp.MustCompile(".*")`,
				Pattern:      `FindAllStringIndex($0, 1)`,
				Returns:      1,
			},
		},
		"FindAllSubmatchIndex": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*regexp.Regexp).FindAllStringSubmatchIndex",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Method: "FindAllStringSubmatchIndex",
			},
			Generate: &checker.Generate{
				PreCondition: `re := regexp.MustCompile(".*")`,
				Pattern:      `FindAllSubmatchIndex($0, 1)`,
				Returns:      1,
			},
		}, //
		"FindAllStringSubmatchIndex": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*regexp.Regexp).FindAllSubmatchIndex",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Method: "FindAllSubmatchIndex",
			},
			Generate: &checker.Generate{
				PreCondition: `re := regexp.MustCompile(".*")`,
				Pattern:      `FindAllStringSubmatchIndex($0, 1)`,
				Returns:      1,
			},
		},
		"FindIndex": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*regexp.Regexp).FindStringIndex",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Method: "FindStringIndex",
			},
			Generate: &checker.Generate{
				PreCondition: `re := regexp.MustCompile(".*")`,
				Pattern:      `FindIndex($0)`,
				Returns:      1,
			},
		},
		"FindStringIndex": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*regexp.Regexp).FindStringIndex",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Method: "FindStringIndex",
			},
			Generate: &checker.Generate{
				PreCondition: `re := regexp.MustCompile(".*")`,
				Pattern:      `FindStringIndex($0)`,
				Returns:      1,
			},
		},
		"FindSubmatchIndex": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*regexp.Regexp).FindStringSubmatchIndex",
			Args:           []int{0},
			StringTargeted: false,
			Alternative: checker.Alternative{
				Method: "FindStringSubmatchIndex",
			},
			Generate: &checker.Generate{
				PreCondition: `re := regexp.MustCompile(".*")`,
				Pattern:      `FindSubmatchIndex($0)`,
				Returns:      1,
			},
		},
		"FindStringSubmatchIndex": {
			Type:           checker.Method,
			Message:        "avoid allocations with (*regexp.Regexp).FindSubmatchIndex",
			Args:           []int{0},
			StringTargeted: true,
			Alternative: checker.Alternative{
				Method: "FindSubmatchIndex",
			},
			Generate: &checker.Generate{
				PreCondition: `re := regexp.MustCompile(".*")`,
				Pattern:      `FindStringSubmatchIndex($0)`,
				Returns:      1,
			},
		},
	}
)
