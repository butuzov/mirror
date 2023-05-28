package checker

import (
	"go/ast"
	"go/importer"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/ast/inspector"
)

func TestViolation(t *testing.T) {
	tests := []struct {
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
			fix := string(v2.suggest(fset))
			if test.ExpectedSuggest != fix {
				t.Errorf("Fix not match: want(%s) vs got(%s)", test.ExpectedSuggest, fix)
			}
			message := v2.Message()
			if test.Message != message {
				t.Errorf("Message not match: want(%s) vs got(%s)", test.Message, message)
			}
		})
	}
}

func TestComplex(t *testing.T) {
	tests := []struct {
		Name            string
		TxtAr           string
		ImportPath      string
		Violation       Violation
		ExpectedMessage string
		ExpectedFix     string
	}{
		{
			Name:       "Simple Check",
			TxtAr:      "testdata/violations_simple.txtar",
			ImportPath: "unicode/utf8",
			Violation: Violation{
				Type:      Function,
				Targets:   Bytes,
				Package:   "unicode/utf8",
				Caller:    "ValidString",
				AltCaller: "Valid",
				Args:      []int{0},
			},
			ExpectedMessage: `avoid allocations with utf8.Valid`,
			ExpectedFix:     `utf8.Valid("foo")`,
		},
		{
			Name:       "Multiline Fix",
			TxtAr:      "testdata/violations_new_lines.txtar",
			ImportPath: "unicode/utf8",
			Violation: Violation{
				Type:      Function,
				Targets:   Bytes,
				Package:   "unicode/utf8",
				Caller:    "ValidString",
				AltCaller: "Valid",
				Args:      []int{0},
			},
			ExpectedMessage: `avoid allocations with utf8.Valid`,
			ExpectedFix:     ``,
		},
		{
			Name:       "Different Packages (No import)",
			TxtAr:      "testdata/violations_packages.txtar",
			ImportPath: "unicode/utf8",
			Violation: Violation{
				Type:       Function,
				Targets:    Bytes,
				Package:    "unicode/utf8",
				AltPackage: "unicode/utf8v2",
				Caller:     "ValidString",
				AltCaller:  "Valid",
				Args:       []int{0},
			},
			ExpectedMessage: `avoid allocations with utf8v2.Valid`,
			ExpectedFix:     ``,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.Name, func(t *testing.T) {
			fset := token.NewFileSet()
			ar, err := Txtar(t, fset, test.TxtAr)
			if err != nil {
				t.Errorf("nil err expected - got %s", err)
			}

			var (
				ins  = inspector.New(ar)
				conf = types.Config{
					Importer: importer.ForCompiler(fset, "source", nil),
				}
				info = types.Info{
					Types: make(map[ast.Expr]types.TypeAndValue),
					Defs:  make(map[*ast.Ident]types.Object),
					Uses:  make(map[*ast.Ident]types.Object),
				}
			)

			// ------ Setup ----------------------------------------------------------

			_, err = conf.Check("source", fset, ar, &info)
			if err != nil {
				t.Errorf("nil err expected - got %s", err)
			}

			check := New([]Violation{test.Violation})
			check.Type = WrapType(&info)
			check.Print = WrapPrint(fset)

			var happend bool

			ins.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
				// allow to check only first call
				if happend {
					return
				}
				happend = true
				// ---- test --------------------------------------------------
				callExpr := n.(*ast.CallExpr)
				expr := callExpr.Fun.(*ast.SelectorExpr)
				x, ok := expr.X.(*ast.Ident)
				if !ok {
					t.Errorf("Expected ast.Indent")
				}

				name := expr.Sel.Name
				// skipping import checks with correct import path
				v := check.Match(test.ImportPath, name)
				if v == nil {
					t.Error("nil violation not expected")
				}

				args, found := check.Handle(v, callExpr)
				if !found {
					t.Errorf("no string to string conversions found")
				}
				v2 := v.With(check.Print(x), callExpr, args)

				gciIssue := v2.Issue(fset)

				if test.ExpectedFix != gciIssue.InlineFix {
					t.Errorf("Fix not match: want(%s) vs got(%s)", test.ExpectedFix, gciIssue.InlineFix)
				}

				if test.ExpectedMessage != gciIssue.Message {
					t.Errorf("Message not match: want(%s) vs got(%s)", test.ExpectedMessage, gciIssue.Message)
				}
			})

			if !happend {
				t.Errorf("Test Not Happend")
			}
		})
	}
}

func TestArgType(t *testing.T) {
	tests := []struct {
		Name      string
		Violation Violation
		Expected  string
	}{
		{
			Name: "DefaultStrings",
			Violation: Violation{
				Targets: Bytes,
			},
			Expected: Strings,
		},
		{
			Name: "DefaultBytes",
			Violation: Violation{
				Targets: Strings,
			},
			Expected: Bytes,
		},
		{
			Name: "NonDefault_Runes",
			Violation: Violation{
				Targets:  Strings,
				ArgsType: Rune,
			},
			Expected: Rune,
		},
		{
			Name: "NonDefault_Byte",
			Violation: Violation{
				Targets:  Strings,
				ArgsType: Byte,
			},
			Expected: Byte,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.Name, func(t *testing.T) {
			if test.Violation.getArgType() != test.Expected {
				t.Errorf("unexpected arg type: want(%s) vs got(%s)", test.Expected, test.Violation.getArgType())
			}
		})
	}
}
