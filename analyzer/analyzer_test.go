package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestCalc(t *testing.T) {
	cases := map[string]struct {
		inputFile string
		expected  []int
	}{
		"not complex": {"./testdata/no-complex.go", []int{1}},
		"if":          {"./testdata/if.go", []int{2}},
		"for":         {"./testdata/for.go", []int{2}},
		"switch":      {"./testdata/switch.go", []int{2}},
	}

	for k, c := range cases {
		t.Run(k, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, c.inputFile, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
			for i, d := range f.Decls {
				fd, ok := d.(*ast.FuncDecl)
				if !ok {
					t.Fatal("decl is not FuncDecl")
				}
				n, err := Calc(fd)
				if err != nil {
					t.Fatal(err)
				}
				if n != c.expected[i] {
					t.Error("Complexity is not expected")
				}
			}
		})
	}
}
