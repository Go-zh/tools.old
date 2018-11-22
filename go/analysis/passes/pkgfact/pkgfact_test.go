package pkgfact_test

import (
	"testing"

	"github.com/Go-zh/tools/go/analysis/analysistest"
	"github.com/Go-zh/tools/go/analysis/passes/pkgfact"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, pkgfact.Analyzer,
		"c", // loads testdata/src/c/c.go.
	)
}
