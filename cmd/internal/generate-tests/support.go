package main

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/butuzov/mirror/internal/rules"
)

type TestCase struct {
	Arguments []string
	Returns   string
	Package   string
	PreCond   string
	Func      string
	Want      string
}

// variate will create a set of variables to use within test fucntion
// if its a string function/method
// -> correct: all strins expected
//
// -> incorrect -> bytes ex
func variate(variance string, oneIsString bool) []string {
	var vars []string
	prefix := "arg"

	for i := 0; i < len(variance); i++ {
		var input string

		// if out argument is in string we gona take argument that is
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
		if strings.ContainsAny(string(s[i]), "().*\\") {
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

// possible varialtions for arguments
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

func makeFunc(pattern string, replaces []string) string {
	for i := 0; i < len(replaces); i++ {
		pattern = strings.Replace(pattern,
			fmt.Sprintf("$%d", i), fmt.Sprintf("arg%d", i), 1)
	}
	return pattern
}

func makeFuncInline(pattern string, replaces []string) string {
	for i := 0; i < len(replaces); i++ {

		variables := strings.Split(replaces[i], ":=")

		pattern = strings.Replace(pattern,
			fmt.Sprintf("$%d", i), strings.Trim(variables[1], " "), 1)
	}
	return pattern
}

func generateTests(pkgName string, list map[string]rules.Diagnostic) []string {
	var tests []string

	// order of the functions
	keys := make([]string, 0, len(list))
	for k := range list {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// tests variance
	for _, key := range keys {
		test := list[key]
		// stop condition with log error
		if test.GenPattern == "" {
			log.Printf("required fields not supported: %#v\n", test)
			continue
		}

		for _, pkg := range PkgImports(pkgName) {
			for _, variance := range PossibleVariations(len(test.Args)) {

				wantStr := ""
				if !strings.Contains(variance, "0") {
					wantStr = QuoteRegexp(test.Message)
				}

				var b1 bytes.Buffer

				// simple test with input variables
				// DONE(butuzov): extenal variables
				// templates.ExecuteTemplate(&b1, "case.tmpl", TestCase{
				// 	Arguments: variate(variance, test.GenString),
				// 	Returns:   GenReturnelements(test.GenReturns),
				// 	Package:   pkg,
				// 	Func:      makeFunc(test.GenPattern, variate(variance, test.GenString)),
				// 	Want:      wantStr,
				// })
				// tests = append(tests, b1.String())

				// TODO(butuzov)  random variables
				// TODO(butuzov): inline variables
				{

					pkgInTest := pkg
					if test.GenCondition != "" {
						pkgInTest = strings.Trim(strings.Split(test.GenCondition, ":=")[0], " ")
					}

					templates.ExecuteTemplate(&b1, "case.tmpl", TestCase{
						Arguments: []string{},
						Returns:   GenReturnelements(test.GenReturns),
						Package:   pkgInTest,
						PreCond:   test.GenCondition,
						Func:      makeFuncInline(test.GenPattern, variate(variance, test.TargetStrings)),
						Want:      wantStr,
					})
					tests = append(tests, b1.String())
					b1.Reset()
				}

			}
		}
	}

	return tests
}
