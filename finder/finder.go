package finder

import "go/ast"

func FindFunc(n *ast.File, f func(*ast.FuncDecl) error) error {
	var err error
	ast.Inspect(n, func(node ast.Node) bool {
		if err != nil {
			return false
		}
		if n, ok := node.(*ast.FuncDecl); ok {
			err = f(n)
			if err != nil {
				return false
			}
		}
		return true
	})
	return err
}
