package finder

import "go/ast"

func FindFuncDecl(n *ast.File, f func(*ast.FuncDecl) error) error {
	var err error
	ast.Inspect(n, func(node ast.Node) bool {
		if err != nil {
			return false
		}
		switch n := node.(type) {
		case *ast.FuncDecl:
			err = f(n)
			if err != nil {
				return false
			}
		}
		return true
	})
	return err
}

func FindFuncLiteral(n *ast.File, f func(*ast.FuncLit) error) error {
	var err error
	ast.Inspect(n, func(node ast.Node) bool {
		if err != nil {
			return false
		}
		switch n := node.(type) {
		case *ast.FuncLit:
			err = f(n)
			if err != nil {
				return false
			}
		}
		return true
	})
	return err
}
