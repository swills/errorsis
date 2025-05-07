package main

import (
        "golang.org/x/tools/go/analysis"

        "errorsis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{errorsis.NoErrorIsStruct}, nil
}
