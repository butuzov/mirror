package mirror

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

// --- debug -------------------------------------------------------------------

func debug(f *token.FileSet) func(ast.Expr, string, ...any) {
	return func(node ast.Expr, format string, a ...any) {
		printer.Fprint(os.Stderr, f, node)
		fmt.Fprintln(os.Stderr, "\n"+strings.Repeat("^", 80))
		fmt.Fprintf(os.Stderr, format, a...)
	}
}

func debugNoOp(_ ast.Expr, _ string, _ ...any) {}
