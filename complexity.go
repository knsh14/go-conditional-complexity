package complexity

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/knsh14/go-conditional-complexity/analyzer"
	"github.com/knsh14/go-conditional-complexity/finder"
	"github.com/knsh14/go-conditional-complexity/message"
)

func Check(path string, threshold int) ([]*message.Message, error) {
	var messages []*message.Message
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}
	finder.FindFunc(f, func(fd *ast.FuncDecl) error {
		count, err := analyzer.Calc(fd)
		if err != nil {
			return err
		}
		if count >= threshold {
			m := message.New(fset, path, fd, count)
			messages = append(messages, m)
		}
		return nil
	})
	return messages, nil
}
