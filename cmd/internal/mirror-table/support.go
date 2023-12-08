package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/butuzov/mirror/internal/checker"
)

// clean up (unrequired for sorting) parts of a string func signature.
func cleanSortKey(a string) string {
	return strings.ReplaceAll(strings.ReplaceAll(a, "(*", ""), ") ", ".")
}

// chaining every violations slice into the one queue.
func chaining[T any](slices ...[]T) chan T {
	c := make(chan T)

	go func(c chan T) {
		defer close(c)

		for i := range slices {
			for _, e := range slices[i] {
				c <- e
			}
		}
	}(c)

	return c
}

// cleanNestedPkg removes slashed path from package name, so
// encoding/utf8 becomes utf8.
func cleanNestedPkg(pkg string) string {
	if strings.Contains(pkg, "/") {
		p := strings.Split(pkg, "/")
		pkg = p[len(p)-1]
	}

	return pkg
}

var funcSignatureArgs = regexp.MustCompile(`[a-zA-Z]{1,}\((.*?)\)$`)

func ternary(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}

func formArgs(v checker.Violation, isAlt bool) string {
	a := []string{}

	if v.Generate != nil && v.Generate.Pattern != "" {
		m := funcSignatureArgs.FindStringSubmatch(v.Generate.Pattern)[1]

		for _, i := range strings.Split(m, ", ") {
			switch {
			case strings.HasPrefix(i, "$"):

				varType := ternary(
					!isAlt,
					ternary(v.Targets != checker.Strings, checker.Bytes, checker.Strings),
					ternary(isAlt && v.ArgsType != "", v.ArgsType,
						ternary(v.Targets != checker.Strings, checker.Strings, checker.Bytes)),
				)

				a = append(a, varType)
			case strings.HasPrefix(i, "\""):
				a = append(a, "string")
			case strings.HasPrefix(i, "'"):
				a = append(a, "byte")
			case strings.HasPrefix(i, "rune"):
				a = append(a, "rune")
			case strings.HasPrefix(i, "1"):
				a = append(a, "int")
			case strings.HasPrefix(i, "func"):
				f := strings.Split(i, "{")
				a = append(a, strings.TrimSpace(f[0]))
			default:
				fmt.Println(">", i)
			}
		}
	}

	return strings.Join(a, ", ")
}

// form returens
func formReturns(ret []string) string {
	if len(ret) == 0 {
		return ""
	}

	if len(ret) == 1 {
		return ret[0]
	}

	return "(" + strings.Join(ret, ", ") + ")"
}

// makes a caller signature
func formCaller(v checker.Violation) string {
	pkg := cleanNestedPkg(v.Package)
	var ret string
	if v.Generate != nil {
		ret = formReturns(v.Generate.Returns)
	}

	if v.Struct != "" {
		return fmt.Sprintf("func (*%s.%s) %s(%s) %s", pkg, v.Struct, v.Caller, formArgs(v, false), ret)
	}
	return fmt.Sprintf("func %s.%s(%s) %s", pkg, v.Caller, formArgs(v, false), ret)
}

// makes an alt caller signature
func formAltCaller(v checker.Violation) string {
	pkg := cleanNestedPkg(v.Package)
	if v.AltPackage != "" {
		pkg = cleanNestedPkg(v.AltPackage)
	}
	var ret string
	if v.Generate != nil {
		ret = formReturns(v.Generate.Returns)
	}

	if v.Struct != "" {
		return fmt.Sprintf("func (*%s.%s) %s(%s) %s", pkg, v.Struct, v.AltCaller, formArgs(v, true), ret)
	}
	return fmt.Sprintf("func %s.%s(%s) %s", pkg, v.AltCaller, formArgs(v, true), ret)
}
