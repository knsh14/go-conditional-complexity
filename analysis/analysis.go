package analysis

import (
	"go/ast"
	"os"
	"path/filepath"

	"github.com/knsh14/go-conditional-complexity/complexity"
	"github.com/knsh14/go-conditional-complexity/result"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = `checks conditional complexity

conditional complexity is index for how function is complex and hard to understand.
`

var threshold int

var Analyzer = &analysis.Analyzer{
	Name:     "conditionalcomplexity",
	Doc:      doc,
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func init() {
	Analyzer.Flags.IntVar(&threshold, "max", 0, "threshold complexity to notice")
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
			fn := pass.Fset.PositionFor(n.Pos(), false).Filename
			wd, err := os.Getwd()
			if err != nil {
				pass.Reportf(n.Pos(), "failed to get working dir: %s", err.Error())
				return
			}
			rel, err := filepath.Rel(wd, fn)
			if err != nil {
				pass.Reportf(n.Pos(), "failed to relative path from %s to %s: %s", wd, fn, err.Error())
				return
			}

			m := result.New(pass.Fset, rel, n, count)
			pass.Reportf(n.Pos(), m.String())
		}
	})
	return nil, nil
}
