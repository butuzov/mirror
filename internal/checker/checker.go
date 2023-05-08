package checker

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Checker will perform standart check on package and its methods
type Checker struct {
	Package   string
	Functions map[string]Violation
	Methods   map[string]map[string]Violation

	types   *types.Info    // used for checking types
	fset    *token.FileSet // debug info
	imports []Import       // imports (of current file)
	debug   func(ast.Expr, string, ...any)
}

// New will accept a name for package (like `text/template` or `strings`) and
// returns a pointer to initial checker object.
func New(importedPackage string) *Checker {
	return &Checker{
		Package:   importedPackage,
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

		// TODO(butuzov): add check for the ast.ParenExpr in e.Fun so we can target
		//                the constructions like this
		// Example:
		//       (&maphash.Hash{}).Write([]byte("foobar"))
		//

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

	// any of ce can be a nil

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

		if !c.isConverterCall(call) {
			continue
		}

		if len(call.Args) == 0 {
			continue
		}

		// checking whats argument
		if v.Targets() != c.Type(call.Args[0]) {
			m[i] = call.Args[0]
		}
	}

	return m, len(m) == len(v.Args)
}

// HandleFunction will return Violation for next processing if function/method
// allows to be violated, so we can check its arguments, after confirming that
// we indeed have method from imported package.
func (c *Checker) HandleFunction(pkgName, methodName string) *Violation {
	m, ok := c.Functions[methodName]
	if !ok || !c.isImported(c.Package, pkgName) {
		return nil
	}

	return &m
}

func (c *Checker) HandleMethod(receiver ast.Expr, method string) *Violation {
	if c.types == nil {
		return nil
	}

	tv := c.types.Types[receiver]
	if !tv.IsValue() || tv.Type == nil {
		return nil
	}

	key := cleanAsterisk(tv.Type.String())
	if methods, ok := c.Methods[key]; !ok {
		return nil
	} else if violation, ok := methods[method]; ok {
		return &violation
	}

	return nil
}

func cleanAsterisk(s string) string {
	if strings.HasPrefix(s, "*") {
		return s[1:]
	}

	return s
}

// isImported will check if package exists in provided imports.
func (c *Checker) isImported(pkg, name string) bool {
	if len(c.imports) == 0 {
		return false
	}

	for _, v := range c.imports {
		if v.Pkg == pkg && v.Name == name {
			return true
		}
	}

	return false
}

func (c *Checker) isConverterCall(ce *ast.CallExpr) bool {
	switch ce.Fun.(type) {
	case *ast.ArrayType, *ast.Ident:
		res := c.Type(ce.Fun)

		return res == "[]byte" || res == "string"
	}

	return false
}

func (c *Checker) Type(node ast.Expr) string {
	// Sometimes it gives what it all about... sometimes not.
	if t := c.types.TypeOf(node); t != nil {
		return t.String()
	}

	if tv, ok := c.types.Types[node]; ok {
		return tv.Type.Underlying().String()
	}

	return ""
}

func (c *Checker) With(pass *analysis.Pass, i []Import, debugFn func(ast.Expr, string, ...any)) *Checker {
	// pass *analysis.Pass
	c.fset = pass.Fset
	c.types = pass.TypesInfo
	c.imports = i
	c.debug = debugFn

	return c
}
