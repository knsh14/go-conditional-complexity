package result

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"sort"
)

type byComplexity []*Message

func (bc byComplexity) Len() int           { return len(bc) }
func (bc byComplexity) Swap(i, j int)      { bc[i], bc[j] = bc[j], bc[i] }
func (bc byComplexity) Less(i, j int) bool { return bc[i].complexity > bc[j].complexity }

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

// FilterByComplexity returns more complex functions than threshold
func FilterByComplexity(msgs []*Message, threshold int) []*Message {
	var ms []*Message
	for _, m := range msgs {
		if m.complexity > threshold {
			ms = append(ms, m)
		}
	}
	return ms
}

// FilterMostComplex returns Most N th complex functions
func FilterMostComplex(msgs []*Message, num int) []*Message {
	sort.Sort(byComplexity(msgs))
	if len(msgs) < num {
		return msgs
	}
	return msgs[:num]
}

// Average returns average complexity of input messages
func Average(msgs []*Message) float64 {
	var total float64
	for _, m := range msgs {
		total += float64(m.complexity)
	}
	return total / float64(len(msgs))
}
