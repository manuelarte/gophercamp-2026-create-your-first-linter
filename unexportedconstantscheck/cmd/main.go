package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/manuelarte/gophercamp-2026-create-your-first-linter/unexportedconstantscheck"
)

func main() {
	singlechecker.Main(unexportedconstantscheck.NewAnalyzer())
}
