package checker

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

// Type of violation: can be method or function
type ViolationType int

const (
	Function ViolationType = iota + 1
	Method
)

// Violation describes what message we going to give to a particular code violation
type Violation struct {
	Type    ViolationType // What type is violation? Method or Function?
	Message string        // Message on violation detection
	Args    []int         // Indexes of the arguments needs to be checked

	StringTargeted bool        // String is expected? []byte otherwise.
	Alternative    Alternative // Alternative methods/functions to use.
	Generate       *Generate   // Rules for generation of tests.
}

type Alternative struct {
	Package  string
	Function string
	Method   string
}

// Tests (generation) related struct.
type Generate struct {
	PreCondition string // Precondition we want to be generated
	Pattern      string // Generate pattern (for the `want` message)
	Returns      int    // Expected to return n elements
}

func (v *Violation) Diagnostic(start, end token.Pos) *analysis.Diagnostic {
	return &analysis.Diagnostic{
		Pos:     start,
		End:     end,
		Message: v.Message,
	}
}

func (v *Violation) Handle(ce *ast.CallExpr) (m map[int]ast.Expr, ok bool) {
	return m, len(m) == len(v.Args)
}

func (v *Violation) Targets() string {
	if !v.StringTargeted {
		return "[]byte"
	}

	return "string"
}

// TODO: not implemented
func (v *Violation) WithAltArgs(m map[int]ast.Expr) *Violation {
	// v.alternativeArgs = m
	return v
}
