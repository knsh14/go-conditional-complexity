package inspector

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/knsh14/go-conditional-complexity/complexity"
	"github.com/knsh14/go-conditional-complexity/finder"
	"github.com/knsh14/go-conditional-complexity/result"
)

// Run returns message of function
func Run(path string) ([]*result.Score, error) {
	var messages []*result.Score
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}
	finder.FindFunc(f, func(fn ast.Node) error {
		count, err := complexity.Count(fn)
		if err != nil {
			return err
		}
		m := result.New(fset, path, fn, count)
		messages = append(messages, m)
		return nil
	})
	return messages, nil
}
