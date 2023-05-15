package checker

import (
	"go/token"
	"testing"

	"golang.org/x/tools/go/ast/inspector"
)

func TestImports(t *testing.T) {
	testData := []struct {
		txtarPath  string
		importsLen int
		hasImports map[string]string
	}{
		{
			txtarPath:  "testdata/imports_nothing.txtar",
			importsLen: 0,
		},
		{
			txtarPath:  "testdata/imports_dot.txtar",
			importsLen: 1,
			hasImports: map[string]string{".": "strings"},
		},
		{
			txtarPath:  "testdata/imports_alias.txtar",
			importsLen: 1,
			hasImports: map[string]string{"foo": "strings"},
		},
		{
			txtarPath:  "testdata/imports_regular.txtar",
			importsLen: 1,
			hasImports: map[string]string{"strings": "strings"},
		},
		{
			txtarPath:  "testdata/imports_all.txtar",
			importsLen: 22,
			hasImports: map[string]string{
				"strings": "strings",
				".":       "strings",
				"foo1":    "strings",
				"foo2":    "strings",
				"foo3":    "strings",
				"foo4":    "strings",
				"foo5":    "strings",
				"foo6":    "strings",
				"foo7":    "strings",
				"foo8":    "strings",
				"foo9":    "strings",
				"foo10":   "strings",
				"foo11":   "strings",
				"foo12":   "strings",
				"foo13":   "strings",
				"foo14":   "strings",
				"foo15":   "strings",
				"foo16":   "strings",
				"foo17":   "strings",
				"foo18":   "strings",
				"foo19":   "strings",
				"foo20":   "strings",
			},
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.txtarPath, func(t *testing.T) {
			t.Parallel()

			fset := token.NewFileSet()
			ar, err := Txtar(t, fset, test.txtarPath)
			if err != nil {
				t.Errorf("nil err expected - got %s", err)
			}

			if len(ar) != 1 {
				t.Errorf("Files in txtar: got(%d) vs want(%d)", len(ar), 1)
			}

			ins := inspector.New(ar)
			testImports := Load(fset, ins)

			// assert
			if len(testImports["a.go"]) != test.importsLen {
				t.Errorf("Imports len not match: got(%d) vs want(%d)", len(testImports["a.go"]), test.importsLen)
			}

			for k, v := range test.hasImports {
				str, ok := testImports.Lookup("a.go", k)
				if !ok {
					t.Errorf("Import `%s` not found", k)
				}

				if v != str {
					t.Errorf("Wrong package found want(%s) vs got(%s)", v, str)
				}
			}

			// test if lookup produce fail
			str, ok := testImports.Lookup("a.go", "foobar")
			if ok {
				t.Errorf("found enexpected package %s", str)
			}
		})
	}
}
