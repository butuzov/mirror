package checker

import (
	"go/ast"
	"go/token"
	"sort"
	"strings"
	"sync"

	"golang.org/x/tools/go/ast/inspector"
)

type Import struct {
	Key string // package name
	Val string // alias
}

type Imports map[string][]Import

// we are going to have Imports entries to be sorted, but if it has less then
// `sortLowerLimit` elements we are skipping this step as its not going to
// be worth of effort.
const sortLowerLimit int = 13

// Package level lock is to prevent import map corruption
var lock sync.RWMutex

func LoadImports(fs *token.FileSet, ins *inspector.Inspector) Imports {
	lock.Lock()
	defer lock.Unlock()

	imports := make(Imports)

	// Populate imports map
	ins.Preorder([]ast.Node{(*ast.ImportSpec)(nil)}, func(node ast.Node) {
		is, _ := node.(*ast.ImportSpec)

		var (
			key   = fs.Position(node.Pos()).Filename
			name  = strings.Trim(is.Path.Value, `"`)
			alias = is.Name.String()
		)

		if is.Name == nil {
			alias = name
		}

		imports[key] = append(imports[key], Import{
			Key: name,
			Val: alias,
		})
	})

	imports.Sort()

	return imports
}

func (i *Imports) Sort() {
	for k := range *i {
		if len((*i)[k]) < sortLowerLimit {
			continue
		}

		k := k
		sort.Slice((*i)[k], func(left, right int) bool {
			return (*i)[k][left].Val < (*i)[k][right].Val
		})
	}
}

func (i Imports) LookupImports(file string) []Import {
	if v, ok := i[file]; ok {
		return v
	}

	return nil
}
