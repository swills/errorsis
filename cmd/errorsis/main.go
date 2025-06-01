package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/swills/errorsis"
)

func main() {
	singlechecker.Main(errorsis.NoErrorIsStruct)
}
