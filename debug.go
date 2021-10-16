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

func debug(fs *token.FileSet, ti map[ast.Expr]types.TypeAndValue) func(ast.Expr) {
	return func(node ast.Expr) {
		tv := ti[node]
		Errorf("how to work with %T of type %#v ?\n", node, tv)

		fmt.Print("affected line> ")
		printer.Fprint(os.Stdout, fs, node)
		fmt.Println()
		fmt.Println(strings.Repeat("^", 80))
	}
}

func debugNoOp(_ ast.Expr) {}

// Errorf is function that reports error
const (
	prefixErr  = "\033[91m [ERR] \033[0m"
	prefixWarn = "\033[93m [ERR] \033[0m"
	prefixInfo = "\033[94m [ERR] \033[0m"
)

func Errorf(format string, args ...interface{}) {
	fmt.Printf(prefixErr+format, args...)
}

func Warnf(format string, args ...interface{}) {
	fmt.Printf(prefixWarn+format, args...)
}

func Infof(format string, args ...interface{}) {
	fmt.Printf(prefixInfo+format, args...)
}
