package rules

var (
	StringFunctions = map[string]Diagnostic{
		"Compare": {
			Message: "this call can be optimized with bytes.Compare",
			Args:    []int{0, 1},

			TargetStrings: true,
			GenPattern:    `Compare($0,$1)`,
			GenReturns:    1,
		},
		"Contains": {
			Message: "this call can be optimized with bytes.Contains",
			Args:    []int{0, 1},

			TargetStrings: true,
			GenPattern:    `Contains($0,$1)`,
			GenReturns:    1,
		},
		"ContainsAny": {
			Message:       "this call can be optimized with bytes.ContainsAny",
			Args:          []int{0},
			TargetStrings: true,
			GenPattern:    `ContainsAny($0,"foobar")`,
			GenReturns:    1,
		},
		"ContainsRune": {
			Message: "this call can be optimized with bytes.ContainsRune",
			Args:    []int{0},

			TargetStrings: true,
			GenPattern:    `ContainsRune($0,'ф')`,
			GenReturns:    1,
		},
		"Count": {
			Message: "this call can be optimized with bytes.Count",
			Args:    []int{0, 1},

			TargetStrings: true,
			GenPattern:    `Count($0, $1)`,
			GenReturns:    1,
		},
		"EqualFold": {
			Message: "this call can be optimized with bytes.EqualFold",
			Args:    []int{0, 1},

			TargetStrings: true,
			GenPattern:    `EqualFold($0,$1)`,
			GenReturns:    1,
		},
		"HasPrefix": {
			Message: "this call can be optimized with bytes.HasPrefix",
			Args:    []int{0, 1},

			TargetStrings: true,
			GenPattern:    `HasPrefix($0,$1)`,
			GenReturns:    1,
		},
		"HasSuffix": {
			Message: "this call can be optimized with bytes.HasSuffix",
			Args:    []int{0, 1},

			TargetStrings: true,
			GenPattern:    `HasSuffix($0,$1)`,
			GenReturns:    1,
		},
		"Index": {
			Message: "this call can be optimized with bytes.Index",
			Args:    []int{0, 1},

			TargetStrings: true,
			GenPattern:    `Index($0,$1)`,
			GenReturns:    1,
		},
		"IndexAny": {
			Message: "this call can be optimized with bytes.IndexAny",
			Args:    []int{0},

			TargetStrings: true,
			GenPattern:    `IndexAny($0, "f")`,
			GenReturns:    1,
		},
		"IndexByte": {
			Message: "this call can be optimized with bytes.IndexByte",
			Args:    []int{0},

			TargetStrings: true,
			GenPattern:    `IndexByte($0, byte('f'))`,
			GenReturns:    1,
		},
		"IndexFunc": {
			Message: "this call can be optimized with bytes.IndexFunc",
			Args:    []int{0},

			TargetStrings: true,
			GenPattern:    `IndexFunc($0,func(r rune) bool { return true })`,
			GenReturns:    1,
		},
		"IndexRune": {
			Message: "this call can be optimized with bytes.IndexRune",
			Args:    []int{0},

			TargetStrings: true,
			GenPattern:    `IndexRune($0, rune('ф'))`,
			GenReturns:    1,
		},
		"LastIndex": {
			Message: "this call can be optimized with bytes.LastIndex",
			Args:    []int{0, 1},

			TargetStrings: true,
			GenPattern:    `LastIndex($0,$1)`,
			GenReturns:    1,
		},
		"LastIndexAny": {
			Message:       "this call can be optimized with bytes.LastIndexAny",
			Args:          []int{0},
			TargetStrings: true,
			GenPattern:    `LastIndexAny($0,"f")`,
			GenReturns:    1,
		},
		"LastIndexByte": {
			Message:       "this call can be optimized with bytes.LastIndexByte",
			Args:          []int{0},
			TargetStrings: true,
			GenPattern:    `LastIndexByte($0, byte('f'))`,
			GenReturns:    1,
		},
		"LastIndexFunc": {
			Message:       "this call can be optimized with bytes.LastIndexAny",
			Args:          []int{0},
			TargetStrings: true,
			GenPattern:    `LastIndexFunc($0, func(r rune) bool { return true })`,
			GenReturns:    1,
		},
	}

	// note(butuzov): adding confidiance feature (flag and field) will allow to check other methods.
	StringsBuilderMethods = map[string]Diagnostic{
		"Write": {
			Message:       "this call can be optimized with WriteString method",
			Args:          []int{0},
			GenCondition:  `builder := strings.Builder{}`,
			TargetStrings: false,
			GenPattern:    `Write($0)`,
			GenReturns:    2,
		},
		"WriteString": {
			Message:       "this call can be optimized with Write method",
			Args:          []int{0},
			GenCondition:  `builder := strings.Builder{}`,
			TargetStrings: true,
			GenPattern:    `WriteString($0)`,
			GenReturns:    2,
		}, //
	}
)
