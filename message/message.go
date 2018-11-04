package message

import (
	"fmt"
	"go/ast"
	"go/token"
)

type Message struct {
	fset       *token.FileSet
	path       string
	name       string
	position   token.Pos
	complexity int
}

func (m *Message) String() string {
	return fmt.Sprintf("%s:%d %s complexity=%d\n", m.path, m.fset.Position(m.position).Line, m.name, m.complexity)
}

func New(fset *token.FileSet, p string, n *ast.FuncDecl, c int) *Message {
	return &Message{
		fset:       fset,
		path:       p,
		name:       n.Name.Name,
		position:   n.Pos(),
		complexity: c,
	}
}
