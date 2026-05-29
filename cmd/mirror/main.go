package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/butuzov/mirror"
)

func main() {
	singlechecker.Main(mirror.NewAnalyzer())
}
