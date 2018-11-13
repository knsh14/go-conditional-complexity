package finder

import (
	"go/ast"

	"github.com/pkg/errors"
)

// FindFunc apply passed function to function node.
func FindFunc(n *ast.File, f func(ast.Node) error) error {
	var err error
	ast.Inspect(n, func(node ast.Node) bool {
		if err != nil {
			return false
		}
		switch n := node.(type) {
		case *ast.FuncDecl, *ast.FuncLit:
			err = f(n)
			if err != nil {
				return false
			}
		}
		return true
	})
	return errors.Wrap(err, "failed to find function")
}
