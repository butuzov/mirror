package checker

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViolation(t *testing.T) {
	for i := range tests {
		test := tests[i]
		t.Run(test.Name, func(t *testing.T) {
			fset := token.NewFileSet()
			expr := ParseExprFrom(t, fset, test.Expression).(*ast.CallExpr)
			args := make(map[int]ast.Expr, len(test.Args))
			base := []byte(test.Base)

			for n := range test.Args {
				args[n] = ParseExprFrom(t, fset, test.Args[n])
			}

			v2 := test.Violation.With(base, expr, args)
			assert.Equal(t, test.ExpectedSuggest, string(v2.suggest(fset)))
			assert.Equal(t, test.Message, v2.Message())
		})
	}
}

var tests = []struct {
	Name            string
	Violation       Violation
	Expression      string
	Base            string
	Args            map[int]string
	ExpectedSuggest string
	Message         string
}{
	{
		Name: "alt Caller",
		Violation: Violation{
			Targets:    Strings,
			Type:       Function,
			Package:    "regexp",
			Caller:     "MatchString",
			Args:       []int{1},
			AltPackage: "foobar",
			AltCaller:  "Match",
		},
		Expression:      `regexp.MatchString("[0-9]+", []bytes("foo"))`,
		Args:            map[int]string{1: `"foo"`},
		Base:            "regexp",
		ExpectedSuggest: `regexp.Match("[0-9]+", "foo")`,
		Message:         `avoid allocations with foobar.Match`,
	},
	{
		Name: "Has More Args Then WeNeed",
		Violation: Violation{
			Targets:    Strings,
			Type:       Function,
			Package:    "regexp",
			Caller:     "MatchString",
			Args:       []int{1},
			AltPackage: "foobar",
			AltCaller:  "Match",
		},
		Expression:      `regexp.MatchString("[0-9]+", []bytes("foo"))`,
		Args:            map[int]string{1: `"foo"`, 2: `"foo"`, 3: `"foo"`, 4: `"foo"`},
		Base:            "regexp",
		ExpectedSuggest: `regexp.Match("[0-9]+", "foo")`,
		Message:         `avoid allocations with foobar.Match`,
	},
	{
		Name: "Regular Suggestion Work",
		Violation: Violation{
			Targets:   Strings,
			Type:      Function,
			Package:   "regexp",
			Caller:    "MatchString",
			Args:      []int{1},
			AltCaller: "Match",
		},
		Expression:      `regexp.MatchString("[0-9]+", []bytes("foo"))`,
		Args:            map[int]string{1: `"foo"`},
		Base:            "regexp",
		ExpectedSuggest: `regexp.Match("[0-9]+", "foo")`,
		Message:         `avoid allocations with regexp.Match`,
	},
	{
		Name: "Methods",
		Violation: Violation{
			Targets:   Strings,
			Type:      Method,
			Package:   "regexp",
			Struct:    "Regexp",
			Caller:    "MatchString",
			Args:      []int{1},
			AltCaller: "Match",
		},
		Expression:      `re.MatchString("[0-9]+", []bytes("foo"))`,
		Args:            map[int]string{1: `"foo"`},
		Base:            "re",
		ExpectedSuggest: `re.Match("[0-9]+", "foo")`,
		Message:         `avoid allocations with (*regexp.Regexp).Match`,
	},
}
