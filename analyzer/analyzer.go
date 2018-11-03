package analyzer

import "go/ast"

func Calc(decl *ast.FuncDecl) (int, error) {
	count := 1
	ast.Inspect(decl, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.IfStmt:
			count++
		case *ast.ForStmt:
			count++
		case *ast.CaseClause:
			count++
		}
		return true
	})
	return count, nil
}
