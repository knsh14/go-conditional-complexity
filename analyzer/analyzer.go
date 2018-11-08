package analyzer

import (
	"go/ast"
	"go/token"
)

// Calc counts conditional complexity of node
func Calc(root ast.Node) (int, error) {
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
