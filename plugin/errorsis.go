package main

import (
	"golang.org/x/tools/go/analysis"

	"github.com/swills/errorsis"
)

func New(_ any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{errorsis.NoErrorIsStruct}, nil
}
