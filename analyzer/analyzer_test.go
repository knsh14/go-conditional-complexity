package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestCalcFuncDecl(t *testing.T) {
	cases := map[string]struct {
		inputFile string
		expected  []int
	}{
		"not complex": {"./testdata/no-complex.go", []int{1}},
		"if":          {"./testdata/if.go", []int{2, 3, 3, 3, 2}},
		"for":         {"./testdata/for.go", []int{2}},
		"switch":      {"./testdata/switch.go", []int{2, 3, 3, 3}},
		"select":      {"./testdata/select.go", []int{2}},
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
				n, err := CalcFuncDecl(fd)
				if err != nil {
					t.Fatal(err)
				}
				if n != c.expected[i] {
					t.Errorf("case %s returns unexpected complexity. expected=%d, got=%d", fd.Name.Name, c.expected[i], n)
				}
			}
		})
	}
}

func TestCaclFuncLit(t *testing.T) {
	cases := []struct {
		title    string
		input    string
		expected int
	}{
		{
			title:    "simple",
			input:    `func() int {return 1}`,
			expected: 1,
		},
		{
			title: "if",
			input: `func(n int) bool {
if n > 10 {
	return false
}
return true
}`,
			expected: 2,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			v, err := parser.ParseExpr(c.input)
			if err != nil {
				t.Fatal(err)
			}
			lit, ok := v.(*ast.FuncLit)
			if !ok {
				t.Fatal("input is not function")
			}
			n, err := CalcFuncLit(lit)
			if err != nil {
				t.Fatal(err)
			}
			if n != c.expected {
				t.Errorf("unexpected complexity. expected=%d, got=%d", c.expected, n)
			}
		})
	}
}
