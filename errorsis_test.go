package errorsis_test

import (
	"testing"

	"github.com/swills/errorsis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestErrorIsStruct(t *testing.T) {
	testDataDir := analysistest.TestData()

	t.Parallel()

	analysistest.Run(t, testDataDir, errorsis.NoErrorIsStruct, "./src/errorsisstruct/")
}
