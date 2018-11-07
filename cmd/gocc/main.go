package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/knsh14/go-conditional-complexity"
)

var (
	threshold int
)

func init() {
	flag.IntVar(&threshold, "max", 12, "threshold to notice")
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
	filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		if strings.Contains(path, "testdata") {
			return nil
		}

		msgs, err := complexity.Check(path, threshold)
		if err != nil {
			return err
		}
		for _, m := range msgs {
			fmt.Fprint(os.Stdout, m)
		}
		return nil
	})
}

func usage() {
}
