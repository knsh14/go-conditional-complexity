package main

import (
	"github.com/knsh14/go-conditional-complexity/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		analysis.Analyzer,
	)
}
