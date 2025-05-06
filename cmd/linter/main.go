package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"errorsis/internal/rules"
)

func main() {
	singlechecker.Main(rules.NoErrorIsStruct)
}
