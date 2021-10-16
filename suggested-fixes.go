package mirror

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/butuzov/mirror/internal/checker"
)

type suggest struct {
	fset    *token.FileSet
	call    *ast.CallExpr
	data    *checker.Violation
	imports map[string][]string
}

func (s suggest) Export() analysis.SuggestedFix {
	return analysis.SuggestedFix{
		Message: "Replace function",
		TextEdits: []analysis.TextEdit{
			{
				Pos:     s.call.Pos(),
				End:     s.call.Pos(),
				NewText: s.Bytes(),
			},
		},
	}
}

func (s suggest) Bytes() []byte {
	var suggest bytes.Buffer

	switch s.data.Type {
	case checker.Method:
		se, _ := s.call.Fun.(*ast.SelectorExpr)
		receiver, _ := se.X.(*ast.Ident)

		suggest.WriteString(receiver.String())
		suggest.WriteByte('.')
		suggest.WriteString(s.data.Alternative.Method)

	case checker.Function:

		switch v := s.call.Fun.(type) {
		case *ast.SelectorExpr:
			// named (regular) import
			receiver, _ := v.X.(*ast.Ident)

			name, found := s.reverseImportSearch(receiver.String(), s.data.Alternative.Package)
			if !found {
				fmt.Printf("(0) add import for %s\n", name)
			}

			suggest.WriteString(name)
			suggest.WriteByte('.')
			suggest.WriteString(s.data.Alternative.Function)

		case *ast.Ident:
			// dot import.
			name, found := s.reverseImportSearch(".", s.data.Alternative.Package)
			// if same?
			if !found {
				fmt.Printf("(1) add import for %s\n", name)
			}
			// what to do?
			if name != s.data.Alternative.Package {
				suggest.WriteString(name)
				suggest.WriteByte('.')
			}

			suggest.WriteString(s.data.Alternative.Function)

		}
		// is ti same package?
		// yes, we keep receiver

		// no we replace receiver

		suggest.WriteString(s.data.Alternative.Function)

	default:
		panic("what is that? not implemented")
	}

	fmt.Fprint(&suggest, "(")

	// arguments.
	var args []string
	for index, v := range s.call.Args {
		if e, ok := s.data.AltArgAt(index); ok {
			args = append(args, types.ExprString(e))
			continue
		}
		args = append(args, types.ExprString(v))
	}

	fmt.Fprint(&suggest, strings.Join(args, ", "))
	fmt.Fprint(&suggest, ")")

	return suggest.Bytes()
}

// reverseImportSearch return same name (if such name exists for any of file imports.) or
// full package
func (s suggest) reverseImportSearch(pkgCur, pkgNew string) (name string, exists bool) {
	var pkgCurName string
	// current
	// todo(butuzov): search for package over map should be more effective. maybe switch to slices of key values.
	for pkg, aliases := range s.imports {
		for _, alias := range aliases {
			if alias == pkgCur {
				// cur
				pkgCurName = pkg
				break
			}
		}

		if pkgCurName != "" {
			break
		}
	}

	// left name in place (as this is same package (e.g. ))
	if pkgCurName == pkgNew {
		return pkgCur, true
	}

	if names, ok := s.imports[pkgNew]; ok {
		return names[0], true
	}

	// we need to add import!
	return pkgNew, false
}
