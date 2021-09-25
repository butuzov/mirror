package rules

var (
	RegexpFunctions = map[string]Diagnostic{
		"Match": {
			Message:       "this call can be optimized with regexp.MatchString",
			Args:          []int{1},
			TargetStrings: false,
			GenPattern:    `Match("foo", $0)`,
			GenReturns:    2,
		},
		"MatchString": {
			Message:       "this call can be optimized with regexp.Match",
			Args:          []int{1},
			TargetStrings: true,
			GenPattern:    `MatchString("foo", $0)`,
			GenReturns:    2,
		},
	}

	// As you see we are not using all of the regexp method because
	// nes we missing return concrete types (bytes or strings)
	// which most probably was intentional.

	// note(butuzov): adding confidiance feature (flag and field) will allow to check other methods.
	RegexpRegexpMethods = map[string]Diagnostic{
		"Match": {
			Message:       "this call can be optimized with (*regexp.Regexp).MatchString",
			Args:          []int{0},
			GenCondition:  `re := regexp.MustCompile(".*")`,
			TargetStrings: false,
			GenPattern:    `Match($0)`,
			GenReturns:    1,
		},
		"MatchString": {
			Message:       "this call can be optimized with (*regexp.Regexp).Match",
			Args:          []int{0},
			GenCondition:  `re := regexp.MustCompile(".*")`,
			TargetStrings: true,
			GenPattern:    `MatchString($0)`,
			GenReturns:    1,
		},
		"FindAllIndex": {
			Message:       "this call can be optimized with (*regexp.Regexp).FindAllStringIndex",
			Args:          []int{0},
			GenCondition:  `re := regexp.MustCompile(".*")`,
			TargetStrings: false,
			GenPattern:    `FindAllIndex($0, 1)`,
			GenReturns:    1,
		},
		"FindAllStringIndex": {
			Message:       "this call can be optimized with (*regexp.Regexp).FindAllIndex",
			Args:          []int{0},
			GenCondition:  `re := regexp.MustCompile(".*")`,
			TargetStrings: true,
			GenPattern:    `FindAllStringIndex($0, 1)`,
			GenReturns:    1,
		},
		"FindAllSubmatchIndex": {
			Message:       "this call can be optimized with (*regexp.Regexp).FindAllStringSubmatchIndex",
			Args:          []int{0},
			GenCondition:  `re := regexp.MustCompile(".*")`,
			TargetStrings: false,
			GenPattern:    `FindAllSubmatchIndex($0, 1)`,
			GenReturns:    1,
		}, //
		"FindAllStringSubmatchIndex": {
			Message:       "this call can be optimized with (*regexp.Regexp).FindAllSubmatchIndex",
			Args:          []int{0},
			GenCondition:  `re := regexp.MustCompile(".*")`,
			TargetStrings: true,
			GenPattern:    `FindAllStringSubmatchIndex($0, 1)`,
			GenReturns:    1,
		},
		"FindIndex": {
			Message:       "this call can be optimized with (*regexp.Regexp).FindStringIndex",
			Args:          []int{0},
			GenCondition:  `re := regexp.MustCompile(".*")`,
			TargetStrings: false,
			GenPattern:    `FindIndex($0)`,
			GenReturns:    1,
		},
		"FindStringIndex": {
			Message:       "this call can be optimized with (*regexp.Regexp).FindStringIndex",
			Args:          []int{0},
			GenCondition:  `re := regexp.MustCompile(".*")`,
			TargetStrings: true,
			GenPattern:    `FindStringIndex($0)`,
			GenReturns:    1,
		},
		"FindSubmatchIndex": {
			Message:       "this call can be optimized with (*regexp.Regexp).FindStringSubmatchIndex",
			Args:          []int{0},
			GenCondition:  `re := regexp.MustCompile(".*")`,
			TargetStrings: false,
			GenPattern:    `FindSubmatchIndex($0)`,
			GenReturns:    1,
		},
		"FindStringSubmatchIndex": {
			Message:       "this call can be optimized with (*regexp.Regexp).FindSubmatchIndex",
			Args:          []int{0},
			GenCondition:  `re := regexp.MustCompile(".*")`,
			TargetStrings: true,
			GenPattern:    `FindStringSubmatchIndex($0)`,
			GenReturns:    1,
		},
	}
)
