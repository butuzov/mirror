package checker

import "go/ast"

// Type of violation: can be methodor function
type Type int

const (
	Function Type = iota + 1
	Method
)

// Violation describs what message we going to give to a particular code violation
type Violation struct {
	Type    Type   // What type is violation? Method (not touching receiver) or Function (changing receiver)?
	Message string // Message we produce on violatin detection
	Args    []int  // Indexes of the arguments needs to be checked

	StringTargeted  bool             // Does string expects to be correct argument ?
	Alternative     Alternative      // Defines names for Diagnostic's Suggestions
	alternativeArgs map[int]ast.Expr // arguments unwrapped (so we can crate suggestion)
	Generate        *Generate        // RUles for tests generations
}

func (v *Violation) WithAltArgs(m map[int]ast.Expr) *Violation {
	v.alternativeArgs = m
	return v
}

func (v *Violation) AltArgAt(idx int) (ast.Expr, bool) {
	val, ok := v.alternativeArgs[idx]
	return val, ok
}

// Generate describes how tests should be generated
type Generate struct {
	PreCondition string // Precondition we want to be generated
	Pattern      string // Generate pattern (for the `want` message)
	Returns      int    // Expected to return n elements
}

type Alternative struct {
	Package  string
	Function string
	Method   string
}
