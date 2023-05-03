package mirror

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"os"
	"strings"
)

// --- debug -------------------------------------------------------------------

func debug(f *token.FileSet, t map[ast.Expr]types.TypeAndValue) func(ast.Expr) {
	return func(node ast.Expr) {
		tv := t[node]
		Errorf("how to work with %T of type %#v ?\n", node, tv)

		fmt.Print("affected line> ")
		printer.Fprint(os.Stdout, f, node)
		fmt.Println()
		fmt.Println(strings.Repeat("^", 80))
	}
}

func debugNoOp(_ ast.Expr) {}

func Errorf(format string, args ...interface{}) {
	fmt.Printf("\033[91m [ERR] \033[0m"+format, args...)
}
