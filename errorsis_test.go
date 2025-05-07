package errorsis_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"errorsis"
)

func TestErrorIsStruct(t *testing.T) {
	testDataDir := analysistest.TestData()

	t.Parallel()

	analysistest.Run(t, testDataDir, errorsis.NoErrorIsStruct, "./src/errorsisstruct/")
}
