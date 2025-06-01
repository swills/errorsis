package errorsis_test

import (
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestErrorIsStruct(t *testing.T) {
	var err error

	var newPlugin register.NewPlugin

	var plugin register.LinterPlugin

	var analyzers []*analysis.Analyzer

	testDataDir := analysistest.TestData()

	t.Parallel()

	newPlugin, err = register.GetPlugin("errorsis")
	if err != nil {
		t.Error(err)
	}

	plugin, err = newPlugin(nil)
	if err != nil {
		t.Error(err)
	}

	analyzers, err = plugin.BuildAnalyzers()
	if err != nil {
		t.Error(err)
	}

	analysistest.Run(t, testDataDir, analyzers[0], "./src/errorsisstruct/")
}
