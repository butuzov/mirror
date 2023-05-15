package checker

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
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

			assert.Nil(t, err)
			assert.Len(t, ar, 1)

			ins := inspector.New(ar)
			testImports := Load(fset, ins)

			// assert
			assert.Len(t, testImports["a.go"], test.importsLen)

			for k, v := range test.hasImports {
				str, ok := testImports.Lookup("a.go", k)
				assert.True(t, ok, "Import `%s` not found", k)
				assert.Equal(t, v, str, "Wrong package found want(%s) vs got(%s)", v, str)
			}

			// test if lookup produce fail
			str, ok := testImports.Lookup("a.go", "foobar")
			assert.False(t, ok, "found somethig enexpected %s", str)
		})
	}
}
