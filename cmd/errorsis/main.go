package main

import (
	"github.com/swills/errorsis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(errorsis.NoErrorIsStruct)
}
