package complexity

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/knsh14/go-conditional-complexity/analyzer"
	"github.com/knsh14/go-conditional-complexity/finder"
	"github.com/knsh14/go-conditional-complexity/result"
)

// Check returns message of too complex function
func Check(path string, threshold int) ([]*result.Message, error) {
	var messages []*result.Message
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}
	finder.FindFuncDecl(f, func(fd *ast.FuncDecl) error {
		count, err := analyzer.CalcFuncDecl(fd)
		if err != nil {
			return err
		}
		if count >= threshold {
			m := result.NewFuncDecl(fset, path, fd, count)
			messages = append(messages, m)
		}
		return nil
	})
	finder.FindFuncLiteral(f, func(fl *ast.FuncLit) error {
		count, err := analyzer.CalcFuncLit(fl)
		if err != nil {
			return err
		}
		if count >= threshold {
			m := result.NewFuncLit(fset, path, fl, count)
			messages = append(messages, m)
		}
		return nil
	})
	return messages, nil
}
