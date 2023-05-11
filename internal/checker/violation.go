package checker

import (
	"fmt"
	"go/ast"
	"path"

	"golang.org/x/tools/go/analysis"
)

// Type of violation: can be method or function
type ViolationType int

const (
	Function ViolationType = iota + 1
	Method
)

const (
	Strings string = "string"
	Bytes   string = "[]byte"
)

// Violation describs what message we going to give to a particular code violation
type Violation struct {
	Type ViolationType //
	Args []int         // Indexes of the arguments needs to be checked

	Targets    string
	Package    string
	AltPackage string
	Struct     string
	Caller     string
	AltCaller  string

	// --- tests generation information
	Generate *Generate

	// --- suggestions related info about violation of rules.
	callExpr  *ast.CallExpr
	arguments map[int]ast.Expr
}

// Tests (generation) related struct.
type Generate struct {
	PreCondition string // Precondition we want to be generated
	Pattern      string // Generate pattern (for the `want` message)
	Returns      int    // Expected to return n elements
}

// func (v *Violation) Diagnostic(fset *token.FileSet, n *ast.CallExpr) *analysis.Diagnostic {
// 	diagnostic := &analysis.Diagnostic{
// 		Pos:     n.Pos(),
// 		End:     n.Pos(),
// 		Message: v.Message,
// 	}

// 	if b := v.suggestion(fset, n); len(b) > 0 {
// 		diagnostic.SuggestedFixes = []analysis.SuggestedFix{{
// 			Message:   "",
// 			TextEdits: []analysis.TextEdit{{Pos: n.Pos(), End: n.End(), NewText: b}},
// 		}}

// 		fmt.Println(">>>", string(b))
// 	}

// 	return diagnostic
// }

// // nolint: revive
// func (v *Violation) suggestion(fset *token.FileSet, n *ast.CallExpr) []byte {
// 	var buf bytes.Buffer

// 	changeCall := func(buf *bytes.Buffer, base, alternative string) []byte {
// 		buf.WriteString(base)
// 		buf.WriteString(".")
// 		buf.WriteString(alternative)
// 		buf.WriteString("(")

// 		for i := range n.Args {
// 			if arg, ok := v.arguments.args[i]; ok {
// 				printer.Fprint(buf, fset, arg)
// 			} else {
// 				printer.Fprint(buf, fset, n.Args[i])
// 			}

// 			if i != len(n.Args)-1 {
// 				buf.WriteString(", ")
// 			}
// 		}

// 		buf.WriteString(")")
// 		return buf.Bytes()
// 	}

// 	// is it method call?
// 	if len(v.arguments.obj) > 0 {
// 		return changeCall(&buf, v.arguments.obj, v.Alternative.Method)
// 	}

// 	// Old And New imports (and names) needs to be resolved.

// 	// rest of cases.
// 	return changeCall(&buf, v.arguments.pkg, v.Alternative.Function)
// }

func (v *Violation) With(e *ast.CallExpr, args map[int]ast.Expr) *Violation {
	v2 := (*v)
	v2.callExpr = e
	v2.arguments = args

	return &v2
}

func (v *Violation) Message() string {
	if v.Type == Method {
		return fmt.Sprintf("avoid allocations with (*%s.%s).%s",
			path.Base(v.Package), v.Struct, v.AltCaller)
	}

	pkg := v.Package
	if len(v.AltPackage) > 0 {
		pkg = v.AltPackage
	}

	return fmt.Sprintf("avoid allocations with %s.%s", path.Base(pkg), v.AltCaller)
}

func (v *Violation) Issue() analysis.Diagnostic {
	diagnostic := analysis.Diagnostic{
		Pos:     v.callExpr.Pos(),
		End:     v.callExpr.Pos(),
		Message: v.Message(),
	}

	if v.Type == Method {
		return diagnostic
	}

	// fmt.Println("package", c.Package)
	// fmt.Println("target methods ?", v.Type == Method)
	// fmt.Println("alternative", v.Alt)

	return diagnostic
}
