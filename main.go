package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/knsh14/go-conditional-complexity/analyzer"
	"github.com/knsh14/go-conditional-complexity/finder"
)

var (
	threshold int
	allfunc   bool
)

func init() {
	flag.IntVar(&threshold, "max", 12, "threshold to notice")
	flag.BoolVar(&allfunc, "a", false, "prints all functions complexity")
}

func main() {
	flag.Parse()
	var p string
	if flag.NArg() == 0 {
		wd, err := os.Getwd()
		p = wd
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to get working directory")
			return
		}
	}
	p = flag.Arg(0)

	ap, err := filepath.Abs(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get Absolute path of %s\n", p)
		return
	}
	filepath.Walk(ap, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		if strings.Contains(path, "testdata") {
			return nil
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return err
		}
		finder.FindFunc(f, func(fd *ast.FuncDecl) error {
			count, err := analyzer.Calc(fd)
			if err != nil {
				return err
			}
			if count >= threshold {
				fmt.Fprintf(os.Stdout, "%s is too complex. complexity=%d\n", fd.Name.Name, count)
			} else if allfunc {
				fmt.Fprintf(os.Stdout, "%s complexity=%d\n", fd.Name.Name, count)
			}
			return nil
		})
		return nil
	})
}

func usage() {
}
