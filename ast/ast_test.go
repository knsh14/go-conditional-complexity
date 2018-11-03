package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFunc(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected int
	}{
		"simple": {"./testdata/simple.go", 1},
	}

	for k, c := range cases {
		t.Run(k, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, c.input, nil, 0)
			if err != nil {
				t.Fatal(err)
			}

			count := 0
			err = FindFunc(f, func(n *ast.FuncDecl) error {
				count++
				return nil
			})
			if err != nil {
				t.Fatal(err)
			}
			if count != c.expected {
				t.Errorf("counted func is not expected. expected=%d, got=%d", c.expected, count)
			}
		})
	}
}
