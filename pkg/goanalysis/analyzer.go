package goanalysis

import (
	"flag"

	"golang.org/x/tools/go/analysis"

	"github.com/incu6us/goimports-reviser/v3/reviser"
)

const errMessage = "imports must be formatted"

func NewAnalyzer(flagSet *flag.FlagSet, localPkgPrefixes string, options ...reviser.SourceFileOption) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func run(localPkgPrefixes string, options ...reviser.SourceFileOption) func(pass *analysis.Pass) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil
}
