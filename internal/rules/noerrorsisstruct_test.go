package rules_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"errorsis/internal/rules"
)

func TestErrorIsStruct(t *testing.T) {
	testDataDir := analysistest.TestData()

	t.Parallel()

	analysistest.Run(t, testDataDir, rules.NoErrorIsStruct, "./src/errorsisstruct/")
}
