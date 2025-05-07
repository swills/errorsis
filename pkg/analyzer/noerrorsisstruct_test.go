package analyzer_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"errorsis/pkg/analyzer"
)

func TestErrorIsStruct(t *testing.T) {
	testDataDir := analysistest.TestData()

	t.Parallel()

	analysistest.Run(t, testDataDir, analyzer.NoErrorIsStruct, "./src/errorsisstruct/")
}
