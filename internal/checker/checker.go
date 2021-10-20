package checker

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"os"

	"github.com/butuzov/mirror/internal/imports"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/tools/go/analysis"
)

// Checker will perform standart check on package and its methods
type Checker struct {
	Package   string
	Functions map[string]Violation
	Methods   map[string]map[string]Violation

	debug   func(ast.Expr)
	types   map[ast.Expr]types.TypeAndValue
	imports []imports.KV
	// temporary fields, assigned to checker just before check.
	tFset *token.FileSet
	tPass *analysis.Pass
}

func New(name string) *Checker {
	return &Checker{
		Package:   name,
		Functions: make(map[string]Violation),
		Methods:   make(map[string]map[string]Violation),

		types: make(map[ast.Expr]types.TypeAndValue),
		tPass: nil,
	}
}

// Debug calls internal debug implementation.
func (c *Checker) Debug(n ast.Expr) {
	if c.debug != nil {
		c.debug(n)
	}
}

// WithDebug sets debug function
func (c *Checker) WithDebug(debugFn func(ast.Expr)) *Checker {
	c.debug = debugFn
	return c
}

// WithFunctions is adding set of functions to be checked.
func (c *Checker) WithFunctions(m map[string]Violation) *Checker {
	if m != nil {
		c.Functions = m
	}
	return c
}

// WithFunctions is adding set of methods of particular struct to be checked.
func (c *Checker) WithStructMethods(structName string, m map[string]Violation) *Checker {
	if m != nil {
		c.Methods[structName] = m
	}
	return c
}

//
func (rc *Checker) WithTypes(typesInfo map[ast.Expr]types.TypeAndValue) *Checker {
	rc.types = typesInfo
	return rc
}

func (rc *Checker) WithImports(imports []imports.KV) *Checker {
	rc.imports = imports
	return rc
}

func (c *Checker) Check(e *ast.CallExpr) *Violation {
	switch v := e.Fun.(type) {
	case *ast.SelectorExpr:

		x, ok := v.X.(*ast.Ident)
		if !ok {
			return nil // can't be mached, so can't be checked.
		}

		// does call expression violates diagnostic rule for package function?
		if d := c.HandleFunction(x.Name, v.Sel.Name); d != nil {
			if argsFixed, found := c.handleDiagnostic(d, e); found {
				return d.WithAltArgs(argsFixed)
			}
		}

		// does call expression violates diagnostic rule for package struct method?
		if d := c.HandleMethod(v.X, v.Sel.Name); d != nil {
			if argsFixed, found := c.handleDiagnostic(d, e); found {
				return d.WithAltArgs(argsFixed)
			}
		}

	case *ast.Ident:

		// special case of "." imported package
		if d := c.HandleFunction(".", v.Name); d != nil {
			// does call expression violates diagnostic rule for package function?
			if argsFixed, found := c.handleDiagnostic(d, e); found {
				return d.WithAltArgs(argsFixed)
			}
		}

	}
	return nil
}

func (c *Checker) HandleMethod(receiver ast.Expr, method string) *Violation {
	if c.types == nil || !c.types[receiver].IsValue() {
		return nil
	}
	tv := c.types[receiver]

	if tv.Type == nil {
		// todo(butuzov): logError
		return nil
	}

	name := tv.Type.String()
	// strip pointer asterisk if it's a pointer
	if name[0] == '*' {
		name = name[1:]
	}

	if methods, ok := c.Methods[name]; !ok {
		return nil
	} else if diagnostic, ok := methods[method]; ok {
		return &diagnostic
	}

	return nil
}

func (c *Checker) HandleFunction(pkgName, methodName string) *Violation {
	m, ok := c.Functions[methodName]
	if !ok || !c.hasImport(c.Package, pkgName) {
		return nil
	}

	return &m
}

// handleDiagnostic will return ErrNoIssue if violation not found, and matches found other wise.
func (c *Checker) handleDiagnostic(d *Violation, ce *ast.CallExpr) (m map[int]ast.Expr, ok bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: %#v\n", err)
			fmt.Println("yeap")
			spew.Dump(d)

			printer.Fprint(os.Stdout, c.tFset, ce)
			os.Exit(1)
		}
	}()

	m = map[int]ast.Expr{}

	maxArgsIdx := len(ce.Args) - 1
	for _, i := range d.Args {
		if i > maxArgsIdx {
			continue
		}

		call, ok := ce.Args[i].(*ast.CallExpr)
		if !ok {
			continue
		}

		// this is []byte(foobarString) or string(foobarBytes) ?
		conv, ok := isConverter(call)
		if !ok || !conv.Valid() {
			continue
		}

		// this argument is []byte while we targeting strings
		if d.StringTargeted && c.Type(conv.Expr().Args[0]) == typeByteSlice {
			m[i] = conv.Expr().Args[0]
		}

		// this argument is string while we targeting []byte
		if !d.StringTargeted && c.Type(conv.Expr().Args[0]) == typeString {
			m[i] = conv.Expr().Args[0]
		}
	}

	return m, len(m) == len(d.Args)
}

// hasImport will check if imports we we have imported pkg as alias?
func (c *Checker) hasImport(pkg, alias string) (res bool) {
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

const (
	// --- string constants for string compare.
	nameStr  = "string"
	nameByte = "byte"
)

func consistsOfBytes(a *ast.ArrayType) bool {
	i, ok := a.Elt.(*ast.Ident)
	if !ok {
		return false
	}

	hasByteElemType := i.Name == nameByte
	return hasByteElemType
}

func indentIsString(i *ast.Ident) bool {
	hasStringAsName := i.Name == nameStr
	return hasStringAsName
}

func NewSuggestedFix(fixes ...analysis.TextEdit) analysis.SuggestedFix {
	var fix analysis.SuggestedFix

	fix.TextEdits = append(fix.TextEdits, fixes...)

	return fix
}
