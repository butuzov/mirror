package main

import (
	"github.com/butuzov/mirror"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(mirror.NewAnalyzer())
}
