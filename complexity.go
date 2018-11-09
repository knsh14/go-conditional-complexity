package complexity

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/knsh14/go-conditional-complexity/analyzer"
	"github.com/knsh14/go-conditional-complexity/finder"
	"github.com/knsh14/go-conditional-complexity/result"
)

// Check returns message of function
func Check(path string) ([]*result.Message, error) {
	var messages []*result.Message
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}
	finder.FindFunc(f, func(fn ast.Node) error {
		count, err := analyzer.Calc(fn)
		if err != nil {
			return err
		}
		m := result.New(fset, path, fn, count)
		messages = append(messages, m)
		return nil
	})
	return messages, nil
}
