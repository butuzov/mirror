package mirror

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), NewAnalyzer())
}

func TestDebug(t *testing.T) {
	wd, _ := os.Getwd()
	analysistest.RunWithSuggestedFixes(t, filepath.Join(wd, "testdata_debug"), NewAnalyzer())
}
