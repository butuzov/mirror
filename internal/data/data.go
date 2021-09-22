package data

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/ast/inspector"
)

type Data struct {
	Types   map[ast.Expr]types.TypeAndValue // <- Comes from types pass.TypesIno
	Imports map[string][]string             // <- We create ourselves
}

func (d *Data) MapExistingImports(ins *inspector.Inspector) {
	d.Imports = make(map[string][]string)

	ins.Preorder([]ast.Node{(*ast.ImportSpec)(nil)}, func(node ast.Node) {
		is, _ := node.(*ast.ImportSpec)
		key := strings.Trim(is.Path.Value, `"`)
		name := is.Name.String()

		if is.Name == nil {
			name = key
		}
		d.Imports[key] = append(d.Imports[key], name)
	})
}

func (d *Data) MapTypesInfo(typesInfo map[ast.Expr]types.TypeAndValue) {
	d.Types = typesInfo
}

// HasImport is checking if slice of names for regexp imports has name in it.
func (d *Data) HasImport(pkg, alias string) bool {
	names, ok := d.Imports[pkg]
	if !ok {
		return false
	}

	for i := range names {
		if names[i] == alias {
			return true
		}
	}
	return false
}
