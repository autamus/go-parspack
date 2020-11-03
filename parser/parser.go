package parser

import (
	"fmt"

	"github.com/autamus/go-parspack/pkg"
	"github.com/autamus/go-parspack/scanner"
)

// Parser is the collection responsible for decoding an input string into
// a package struct.
type Parser struct {
	scnr   scanner.Scanner
	result pkg.Package
}

// Init initializes a Parser with a scanner.
func (parser *Parser) Init(scnr scanner.Scanner) {
	parser.scnr = scnr
	parser.result = pkg.Package{}
}

// Parse decodes the information from a scanner into a package struct.
func (parser *Parser) Parse() (result pkg.Package, err error) {
	if !parser.scnr.HasNextLine() && !parser.scnr.HasNextOnLine() {
		return parser.result, nil
	}

	token := parser.scnr.Get()

	switch {
	case token.IsComment() && parser.scnr.HasNextLine():
		parser.scnr.NextLine()
		parser.Parse()

	case token.IsString():
		fmt.Println("Description")

	case token.IsClass():
		err = parser.ParseClass()
		if err != nil {
			return result, err
		}
	}

	err = parser.scnr.Next()
	if err != nil {
		if err.Error() == "end of input string" {
			return parser.result, nil
		}
		return parser.result, err
	}
	return parser.Parse()
}
