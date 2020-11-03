package parspack

import (
	"github.com/autamus/go-parspack/pkg"
	"github.com/autamus/go-parspack/scanner"

	"github.com/autamus/go-parspack/parser"
)

// Decode parses a segment of Spack build instructional syntax
// and returns the results in a package Struct.
func Decode(spack string) (result pkg.Package, err error) {
	scnr := scanner.Scanner{}
	parsr := parser.Parser{}
	scnr.Init(spack)
	parsr.Init(scnr)
	result, err = parsr.Parse()
	if err != nil {
		return result, err
	}
	return result, nil
}
