package checker

import (
	"go/ast"
	"go/types"
)

// Checker will perform standart check on package and its methods
type Checker struct {
	Package   string
	Functions map[string]Violation
	Methods   map[string]map[string]Violation

	fTypes  *types.Info
	imports []Import
}

func New(name string) *Checker {
	return &Checker{
		Package:   name,
		Functions: make(map[string]Violation),
		Methods:   make(map[string]map[string]Violation),
	}
}

// Check perform check on call expression in order to find out can the call
// expression to be substituted with an alternative method/function.
func (c *Checker) Check(e *ast.CallExpr) *Violation {
	switch expr := e.Fun.(type) {

	// Regular calls (`*ast.SelectorExpr`) like strings.HasPrefix or re.Match are
	// handled by this check
	case *ast.SelectorExpr:
		x, ok := expr.X.(*ast.Ident)
		if !ok {
			return nil // can't be mached, so can't be checked.
		}

		// does call expression violates diagnostic rule for package function?
		if v := c.HandleFunction(x.Name, expr.Sel.Name); v != nil {
			if argsFixed, found := c.handleViolation(v, e); found {
				return v.WithAltArgs(argsFixed)
			}
		}

		// does call expression violates diagnostic rule for package struct method?
		if v := c.HandleMethod(expr.X, expr.Sel.Name); v != nil {
			if argsFixed, found := c.handleViolation(v, e); found {
				return v.WithAltArgs(argsFixed)
			}
		}

	// Special case of "." imported packages
	case *ast.Ident:
		// special case of "." imported package
		if v := c.HandleFunction(".", expr.Name); v != nil {
			// does call expression violates diagnostic rule for package function?
			if argsFixed, found := c.handleViolation(v, e); found {
				return v.WithAltArgs(argsFixed)
			}
		}

	}
	return nil
}

func (c *Checker) handleViolation(v *Violation, ce *ast.CallExpr) (map[int]ast.Expr, bool) {
	m := map[int]ast.Expr{}

	for _, i := range v.Args {
		if i >= len(ce.Args) {
			continue
		}

		call, ok := ce.Args[i].(*ast.CallExpr)
		if !ok {
			continue
		}

		if t := c.Type(call); t.String() != "string" && t.String() != "[]byte" {
			continue
		}

		if string(v.Targets()) != c.fTypes.TypeOf(call.Args[0]).String() {
			m[i] = call.Args[0]
		}

	}
	return m, len(m) == len(v.Args)
}

func (c *Checker) Type(e ast.Expr) types.Type {
	return c.fTypes.TypeOf(e)
}

// HandleFunction will return Violation for next processing if function/method
// allows to be violated, so we can check its arguments, after confirming that
// we indeed have method from imported package.
func (c *Checker) HandleFunction(pkgName, methodName string) *Violation {
	m, ok := c.Functions[methodName]
	if !ok || !c.imported(c.Package, pkgName) {
		return nil
	}

	return &m
}

func (c *Checker) HandleMethod(receiver ast.Expr, method string) *Violation {
	if c.fTypes == nil || !c.fTypes.Types[receiver].IsValue() {
		return nil
	}
	tv := c.fTypes.Types[receiver]

	if tv.Type == nil {
		// todo(butuzov): logError
		return nil
	}

	if methods, ok := c.Methods[cleanName(tv.Type.String())]; !ok {
		return nil
	} else if violation, ok := methods[method]; ok {
		return &violation
	}

	return nil
}

// imported will check if package exists in provided imports.
func (c *Checker) imported(pkg, alias string) bool {
	if len(c.imports) == 0 {
		return false
	}

	for _, v := range c.imports {
		if v.Val == alias && v.Key == pkg {
			return true
		}
	}

	return false
}

// cleanName will remove * from the name variable if it is a pointer.
func cleanName(name string) string {
	if name[0] == '*' {
		return name[1:]
	}
	return name
}

// -----------------------------------------------------------------------------
type (
	// This is type aliases only in order to make code a bit more readable.
	typesInfo = map[ast.Expr]types.TypeAndValue
	imports   = []Import
)

func (c *Checker) With(types *types.Info, i imports) *Checker {
	c.fTypes = types
	c.imports = i

	return c
}