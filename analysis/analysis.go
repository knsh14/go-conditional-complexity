package gocc

import (
	"flag"
	"go/ast"

	"github.com/knsh14/go-conditional-complexity/complexity"
	"github.com/knsh14/go-conditional-complexity/result"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = `checks conditional complexity

conditional complexity is index for how function is complex and hard to understand.
`

var (
	FlagSet   *flag.FlagSet
	threshold int
)

func init() {
	FlagSet = flag.NewFlagSet("gocc", flag.ExitOnError)
	FlagSet.IntVar(&threshold, "max", 12, "threshold to notice")
}

var Analyzer = &analysis.Analyzer{
	Name:  "conditional complexity",
	Doc:   doc,
	Flags: *FlagSet,
	Run:   run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		count, err := complexity.Count(n)
		if err != nil {
			pass.Reportf(n.Pos(), "failed to count: %s", err.Error())
			return
		}
		if threshold < count {
			m := result.New(pass.Fset, "", n, count)
			pass.Reportf(n.Pos(), m.String())
		}
	})
	return nil, nil
}
