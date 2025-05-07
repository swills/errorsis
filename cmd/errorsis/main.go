package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"errorsis"
)

func main() {
	singlechecker.Main(errorsis.NoErrorIsStruct)
}
