package result

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
)

// Message to output result
type Message struct {
	fset       *token.FileSet
	path       string
	name       string
	position   token.Pos
	complexity int
}

func (m *Message) String() string {
	return fmt.Sprintf("%s:%d func %s complexity=%d\n", m.path, m.fset.Position(m.position).Line, m.name, m.complexity)
}

// New returns result info for Function
func New(fset *token.FileSet, p string, n ast.Node, c int) *Message {
	funcName := "literal"
	if fd, ok := n.(*ast.FuncDecl); ok {
		var buf bytes.Buffer
		if fd.Recv != nil {
			buf.WriteString("(")
			for _, r := range fd.Recv.List {
				buf.WriteString(r.Names[0].Name)
				buf.WriteString(" ")
				if t, ok := r.Type.(*ast.StarExpr); ok {
					buf.WriteString("*")
					buf.WriteString(t.X.(*ast.Ident).Name)
				}
				if t, ok := r.Type.(*ast.Ident); ok {
					buf.WriteString(t.Name)
				}
			}
			buf.WriteString(")")
		}
		buf.WriteString(" " + fd.Name.Name)
		funcName = buf.String()
	}
	return &Message{
		fset:       fset,
		path:       p,
		name:       funcName,
		position:   n.Pos(),
		complexity: c,
	}
}
