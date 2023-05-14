package checker

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"strings"
	"testing"

	"golang.org/x/tools/txtar"
)

func Txtar(t *testing.T, fset *token.FileSet, txtarPath string) (files []*ast.File, err error) {
	t.Helper()

	ar, err := txtar.ParseFile(txtarPath)
	if err != nil {
		return nil, err
	}

	files = make([]*ast.File, 0, len(ar.Files))
	for i := range ar.Files {
		file := ar.Files[i]
		if !strings.HasSuffix(ar.Files[i].Name, ".go") {
			continue
		}

		f, err := parser.ParseFile(fset,
			path.Base(file.Name), file.Data, parser.AllErrors)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}

	return files, nil
}
