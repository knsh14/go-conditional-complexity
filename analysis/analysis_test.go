package analysis_test

import (
	"testing"

	"github.com/knsh14/go-conditional-complexity/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analysis.Analyzer, "complexity")
}
