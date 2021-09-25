package rules

var (
	BytesFunctions = map[string]Diagnostic{
		"NewBuffer": {
			Message:       "this call can be optimized with bytes.NewBufferString",
			Args:          []int{0},
			TargetStrings: false,
			GenPattern:    `NewBuffer($0)`,
			GenReturns:    1,
		},
		"NewBufferString": {
			Message:       "bytes.NewBuffer",
			Args:          []int{0},
			TargetStrings: true,
			GenPattern:    `NewBufferString($0)`,
			GenReturns:    1,
		},
		"Compare": {
			Message:       "this call can be optimized with strings.Compare",
			Args:          []int{0, 1},
			TargetStrings: false,
			GenPattern:    `Compare($0, $1)`,
			GenReturns:    1,
		},
		"Contains": {
			Message:       "this call can be optimized with strings.Contains",
			Args:          []int{0, 1},
			TargetStrings: false,
			GenPattern:    `Contains($0, $1)`,
			GenReturns:    1,
		},
		"ContainsAny": {
			Message:       "this call can be optimized with strings.ContainsAny",
			Args:          []int{0},
			TargetStrings: false,
			GenPattern:    `ContainsAny($0, "f")`,
			GenReturns:    1,
		},
		"ContainsRune": {
			Message:       "this call can be optimized with strings.ContainsRune",
			Args:          []int{0},
			TargetStrings: false,
			GenPattern:    `ContainsRune($0, rune('ф'))`,
			GenReturns:    1,
		},
		"Count": {
			Message:       "this call can be optimized with strings.Count",
			Args:          []int{0, 1},
			TargetStrings: false,
			GenPattern:    `Count($0, $1)`,
			GenReturns:    1,
		},
		"EqualFold": {
			Message:       "this call can be optimized with strings.EqualFold",
			Args:          []int{0, 1},
			TargetStrings: false,
			GenPattern:    `EqualFold($0, $1)`,
			GenReturns:    1,
		},
		"HasPrefix": {
			Message:       "this call can be optimized with strings.HasPrefix",
			Args:          []int{0, 1},
			TargetStrings: false,
			GenPattern:    `HasPrefix($0, $1)`,
			GenReturns:    1,
		},
		"HasSuffix": {
			Message:       "this call can be optimized with strings.HasSuffix",
			Args:          []int{0, 1},
			TargetStrings: false,
			GenPattern:    `HasSuffix($0, $1)`,
			GenReturns:    1,
		},
		"Index": {
			Message:       "this call can be optimized with strings.Index",
			Args:          []int{0, 1},
			TargetStrings: false,
			GenPattern:    `Index($0, $1)`,
			GenReturns:    1,
		},
		"IndexAny": {
			Message:       "this call can be optimized with strings.IndexAny",
			Args:          []int{0},
			TargetStrings: false,
			GenPattern:    `IndexAny($0, "f")`,
			GenReturns:    1,
		},
		"IndexByte": {
			Message:       "this call can be optimized with strings.IndexByte",
			Args:          []int{0},
			TargetStrings: false,
			GenPattern:    `IndexByte($0, 'f')`,
			GenReturns:    1,
		},
		"IndexFunc": {
			Message:       "this call can be optimized with strings.IndexFunc",
			Args:          []int{0},
			TargetStrings: false,
			GenPattern:    `IndexFunc($0, func(rune) bool {return true })`,
			GenReturns:    1,
		},
		"IndexRune": {
			Message:       "this call can be optimized with strings.IndexRune",
			Args:          []int{0},
			TargetStrings: false,
			GenPattern:    `IndexRune($0, rune('ф'))`,
			GenReturns:    1,
		},
		"LastIndex": {
			Message:       "this call can be optimized with strings.LastIndex",
			Args:          []int{0, 1},
			TargetStrings: false,
			GenPattern:    `LastIndex($0, $1)`,
			GenReturns:    1,
		},
		"LastIndexAny": {
			Message:       "this call can be optimized with strings.LastIndexAny",
			Args:          []int{0},
			TargetStrings: false,
			GenPattern:    `LastIndexAny($0, "ф")`,
			GenReturns:    1,
		},

		"LastIndexByte": {
			Message:       "this call can be optimized with strings.LastIndexByte",
			Args:          []int{0},
			TargetStrings: false,
			GenPattern:    `LastIndexByte($0, 'f')`,
			GenReturns:    1,
		},
		"LastIndexFunc": {
			Message:       "this call can be optimized with strings.LastIndexAny",
			Args:          []int{0},
			TargetStrings: false,
			GenPattern:    `LastIndexFunc($0, func(rune) bool {return true })`,
			GenReturns:    1,
		},
	}

	// note(butuzov): adding confidiance feature (flag and field) will allow to check other methods.
	BytesBufferMethods = map[string]Diagnostic{
		"Write": {
			Message:       "you should be using WriteString method",
			Args:          []int{0},
			TargetStrings: false,
			GenCondition:  "bb := bytes.Buffer{}",
			GenPattern:    `Write($0)`,
			GenReturns:    2,
		},
		"WriteString": {
			Message:       "you should be using Write method",
			Args:          []int{0},
			TargetStrings: true,
			GenCondition:  "bb := bytes.Buffer{}",
			GenPattern:    `WriteString($0)`,
			GenReturns:    2,
		}, //
	}
)
