package result

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"sort"
)

// Score to output result
type Score struct {
	path       string
	line       int
	name       string
	complexity int
}

func (s *Score) String() string {
	return fmt.Sprintf("%s:%d func %s complexity=%d\n", s.path, s.line, s.name, s.complexity)
}

type byComplexity []*Score

func (bc byComplexity) Len() int           { return len(bc) }
func (bc byComplexity) Swap(i, j int)      { bc[i], bc[j] = bc[j], bc[i] }
func (bc byComplexity) Less(i, j int) bool { return bc[i].complexity > bc[j].complexity }

// New returns result info for Function
func New(fset *token.FileSet, p string, n ast.Node, c int) *Score {
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
	return &Score{
		path:       p,
		line:       fset.Position(n.Pos()).Line,
		name:       funcName,
		complexity: c,
	}
}

// FilterByComplexity returns more complex functions than threshold
func FilterByComplexity(msgs []*Score, threshold int) []*Score {
	var ms []*Score
	for _, m := range msgs {
		if m.complexity > threshold {
			ms = append(ms, m)
		}
	}
	return ms
}

// FilterMostComplex returns Most N th complex functions
func FilterMostComplex(msgs []*Score, num int) []*Score {
	sort.Sort(byComplexity(msgs))
	if len(msgs) < num {
		return msgs
	}
	return msgs[:num]
}

// Average returns average complexity of input messages
func Average(msgs []*Score) float64 {
	var total float64
	for _, m := range msgs {
		total += float64(m.complexity)
	}
	return total / float64(len(msgs))
}
