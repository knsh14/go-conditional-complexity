package complexity

import (
	"errors"
	"go/ast"
	"go/token"
)

// Count returns conditional complexity of function node
func Count(root ast.Node) (int, error) {
	switch root.(type) {
	case *ast.FuncDecl, *ast.FuncLit:
		break
	default:
		return 0, errors.New("root node is not function declaration or function literal")
	}
	count := 1
	ast.Inspect(root, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.IfStmt:
			count++
			count += checkAndOrNode(n.Cond)
		case *ast.ForStmt:
			count++
		case *ast.CaseClause:
			// if n.List is nil. it means default clause.
			if n.List == nil {
				count++
				break
			}
			count += len(n.List)
			for _, l := range n.List {
				count += checkAndOrNode(l)
			}
		case *ast.CommClause:
			count++
		case *ast.FuncLit:
			if node != root {
				return false
			}
		}
		return true
	})
	return count, nil
}

func checkAndOrNode(n ast.Node) int {
	count := 0
	ast.Inspect(n, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.BinaryExpr:
			if n.Op == token.LAND || n.Op == token.LOR {
				count++
			}
		}
		return true
	})
	return count
}
