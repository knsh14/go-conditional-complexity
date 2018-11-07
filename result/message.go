package result

import (
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
	return fmt.Sprintf("%s:%d %s complexity=%d\n", m.path, m.fset.Position(m.position).Line, m.name, m.complexity)
}

// NewFuncDecl returns struct instance for output
func NewFuncDecl(fset *token.FileSet, p string, n *ast.FuncDecl, c int) *Message {
	return &Message{
		fset:       fset,
		path:       p,
		name:       n.Name.Name,
		position:   n.Pos(),
		complexity: c,
	}
}

// NewFuncLit returns result info for Function Literal
func NewFuncLit(fset *token.FileSet, p string, n *ast.FuncLit, c int) *Message {
	return &Message{
		fset:       fset,
		path:       p,
		name:       "Literal",
		position:   n.Pos(),
		complexity: c,
	}
}
