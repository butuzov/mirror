package main

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/butuzov/mirror/internal/checker"
)

type TestCase struct {
	Arguments []string
	Returns   string
	Package   string
	PreCond   string
	Func      string
	Want      string
}

// variate will create a set of variables to use within test function
// if its a string function/method
// -> correct: all strings expected
//
// -> incorrect -> bytes ex
func variate(variance string, oneIsString bool) []string {
	var vars []string
	prefix := "arg"

	for i := 0; i < len(variance); i++ {
		var input string

		// if out argument is in string we gonna take argument that is
		if oneIsString {
			input = `"foobar"`
		} else {
			input = `[]byte{'f','o','o','b','a','r'}`
		}

		if oneIsString && (variance[i] == '1') {
			input = fmt.Sprintf(`%s(%s)`, "[]byte", input)
		}
		if !oneIsString && (variance[i] == '0') {
			input = fmt.Sprintf(`%s(%s)`, "string", input)
		}

		if oneIsString {
			// looking for string
			if variance[i] == '1' {
				input = `string([]byte{'f','o','o','b','a','r'})`
			} else {
				input = `"foobar"`
			}
		} else {
			// looking for byte
			if variance[i] == '0' {
				input = `[]byte{'f','o','o','b','a','r'}`
			} else {
				input = `[]byte("foobar")`
			}
		}

		vars = append(vars, fmt.Sprintf("%s%d := %s", prefix, i, input))
	}
	return vars
}

// import, dot-import, alias
func PkgImports(pkg string) []string {
	return []string{pkg, "", "pkg"}
}

// quote regexp

func QuoteRegexp(s string) string {
	var str strings.Builder
	for i := range s {
		if bytes.ContainsAny([]byte{s[i]}, "().*\\") {
			str.WriteByte('\\')
		}
		str.WriteByte(s[i])
	}
	return str.String()
}

func GenReturnelements(n int) string {
	var ret []byte
	for i := 0; i < n; i++ {
		ret = append(ret, '_')
		if i != (n - 1) {
			ret = append(ret, ',')
		}
	}
	return string(ret)
}

// possible variations for arguments
func PossibleVariations(n int) []string {
	if n <= 1 {
		return []string{"1", "0"}
	}

	var out []string
	format := "%0" + fmt.Sprintf("%d", n) + "[1]b"
	for i := 0; i <= all(n); i++ {
		out = append(out, fmt.Sprintf(format, i))
	}
	return out
}

func all(n int) int {
	if n <= 1 {
		return n
	}
	return (1 << (n)) - 1
}

func makeFuncInline(pattern string, replaces []string) string {
	for i := 0; i < len(replaces); i++ {

		variables := strings.Split(replaces[i], ":=")

		pattern = strings.Replace(pattern,
			fmt.Sprintf("$%d", i), strings.Trim(variables[1], " "), 1)
	}
	return pattern
}

func generateTests(pkgName string, list []checker.Violation) []string {
	var tests []string

	sort.Slice(list, func(i, j int) bool {
		return list[i].Caller < list[j].Caller
	})

	// tests variance
	for _, test := range list {

		if test.Generate == nil || test.Generate.Pattern == "" {
			log.Printf("required fields not supported: %#v\n", test)
			continue
		}

		for _, pkg := range PkgImports(pkgName) {
			for _, variance := range PossibleVariations(len(test.Args)) {
				wantStr := ""
				if !strings.Contains(variance, "0") {
					wantStr = QuoteRegexp(test.Message())
				}

				if test.Generate.SkipGenerate {
					continue
				}

				var buf bytes.Buffer

				pkgInTest := pkg
				preCondition := test.Generate.PreCondition
				if test.Generate.PreCondition != "" {
					pkgInTest = strings.Trim(strings.Split(test.Generate.PreCondition, ":=")[0], " ")

					alt := pkg + "."
					if strings.Trim(pkg, " ") == "" {
						alt = ""
					}

					preCondition = strings.Replace(preCondition, pkgName+".", alt, -1)

				}

				templates.ExecuteTemplate(&buf, "case.tmpl", TestCase{
					Arguments: []string{},
					Returns:   GenReturnelements(len(test.Generate.Returns)),
					Package:   pkgInTest,
					PreCond:   preCondition,
					Func: makeFuncInline(test.Generate.Pattern,
						variate(variance, test.Targets == checker.Strings)),
					Want: wantStr,
				})

				tests = append(tests, buf.String())
				buf.Reset()
			}
		}

	}

	return tests
}
