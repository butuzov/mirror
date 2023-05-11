package checker

import (
	"go/ast"
	"strings"
)

// Checker will perform standart check on package and its methods.
type Checker struct {
	Violations []Violation           // List of available violations
	Packages   map[string][]int      // Storing indexes of Violations per pkg/kg.Struct
	Type       func(ast.Expr) string // Closure for the
}

func New(violations ...[]Violation) Checker {
	c := Checker{
		Packages: make(map[string][]int),
	}

	for i := range violations {
		c.register(violations[i])
	}

	return c
}

// Match will check the available violations we got from checks against
// the `name` caller from package `pkgName`.
func (c *Checker) Match(pkgName, name string) *Violation {
	// Does it have struct?
	checkStruct := strings.Contains(pkgName, ".")

	for _, idx := range c.Packages[pkgName] {
		if c.Violations[idx].Caller == name {
			if checkStruct == (len(c.Violations[idx].Struct) == 0) {
				continue
			}

			// copy violation
			v := c.Violations[idx]

			return &v
		}
	}

	return nil
}

func (c *Checker) Handle(v *Violation, ce *ast.CallExpr) (map[int]ast.Expr, bool) {
	m := map[int]ast.Expr{}

	// We going to check each of elements we mark for checking, in order to find,
	// a call that violates our rules.
	for _, i := range v.Args {
		if i >= len(ce.Args) {
			continue
		}

		call, ok := ce.Args[i].(*ast.CallExpr)
		if !ok {
			continue
		}

		// is it convertsion call
		if !c.callConverts(call) {
			continue
		}

		// somehow no argument of call
		if len(call.Args) == 0 {
			continue
		}

		// wrong argument type
		if v.Targets == c.Type(call.Args[0]) {
			continue
		}

		m[i] = call.Args[0]
	}

	return m, len(m) == len(v.Args)
}

func (c *Checker) callConverts(ce *ast.CallExpr) bool {
	switch ce.Fun.(type) {
	case *ast.ArrayType, *ast.Ident:
		res := c.Type(ce.Fun)

		return res == "[]byte" || res == "string"
	}

	return false
}

// register violations.
func (c *Checker) register(violations []Violation) {
	for _, v := range violations { // nolint: gocritic
		c.Violations = append(c.Violations, v)
		if len(v.Struct) > 0 {
			c.registerIdxPer(v.Package + "." + v.Struct)
		}
		c.registerIdxPer(v.Package)
	}
}

// registerIdxPer will register last added violation element
// under pkg string.
func (c *Checker) registerIdxPer(pkg string) {
	c.Packages[pkg] = append(c.Packages[pkg], len(c.Violations)-1)
}
