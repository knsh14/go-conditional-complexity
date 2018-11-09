package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/knsh14/go-conditional-complexity"
	"github.com/knsh14/go-conditional-complexity/result"
)

var (
	threshold int
	exclude   string
	top       int
	avg       bool
)

func init() {
	flag.IntVar(&threshold, "max", 12, "threshold to notice")
	flag.StringVar(&exclude, "exclude", "", "exclude file path pattern")
	flag.IntVar(&top, "top", 0, "show highest complicated functions")
	flag.BoolVar(&avg, "avg", false, "show average of all")
}

func main() {
	flag.Parse()
	p := flag.Arg(0)
	if flag.NArg() == 0 {
		wd, err := os.Getwd()
		p = wd
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to get working directory")
			return
		}
	}

	if _, err := os.Stat(p); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s is not exist\n", p)
		return
	}
	var excludePattern *regexp.Regexp
	if exclude != "" {
		p, err := regexp.Compile(exclude)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return
		}
		excludePattern = p
	}
	var allmessages []*result.Message
	filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		if excludePattern != nil && excludePattern.MatchString(path) {
			return nil
		}

		msgs, err := complexity.Check(path)
		if err != nil {
			return err
		}
		allmessages = append(allmessages, msgs...)
		return nil
	})
	msgs := result.FilterByComplexity(allmessages, threshold)
	if top > 0 {

		for _, m := range result.FilterMostComplex(msgs, top) {
			fmt.Fprint(os.Stdout, m)
		}
	}
	if avg {
		allAvg := result.Average(allmessages)
		filteredAvg := result.Average(msgs)
		fmt.Fprintf(os.Stdout, "All Funcs Average:%f, Complex Funcs Average:%f\n", allAvg, filteredAvg)
	}
}

func usage() {
}
