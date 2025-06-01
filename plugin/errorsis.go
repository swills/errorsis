package main

import (
	"github.com/swills/errorsis"
	"golang.org/x/tools/go/analysis"
)

func New(_ any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{errorsis.NoErrorIsStruct}, nil
}
