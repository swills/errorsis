package main

import (
	"golang.org/x/tools/go/analysis"

	"errorsis"
)

func New(_ any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{errorsis.NoErrorIsStruct}, nil
}
