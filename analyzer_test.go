package mirror

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), NewAnalyzer())
}

func TestDebug(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), NewAnalyzer())
}
