package gocc

import (
	"flag"
	"go/ast"

	"github.com/knsh14/go-conditional-complexity/complexity"
	"github.com/knsh14/go-conditional-complexity/finder"
	"github.com/knsh14/go-conditional-complexity/result"
	"golang.org/x/tools/go/analysis"
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
	for _, f := range pass.Files {
		finder.FindFunc(f, func(fn ast.Node) error {
			count, err := complexity.Count(fn)
			if err != nil {
				return err
			}
			if threshold < count {
				m := result.New(pass.Fset, "", fn, count)
				pass.Reportf(fn.Pos(), m.String())
			}
			return nil
		})
	}
	return nil, nil
}
