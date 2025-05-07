package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"errorsis/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.NoErrorIsStruct)
}
