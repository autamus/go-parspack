package parser

import (
	"github.com/autamus/go-parspack/pkg"
	"github.com/autamus/go-parspack/scanner"
)

// Parser is the collection responsible for decoding an input string into
// a package struct.
type Parser struct {
	scnr   scanner.Scanner
	result *pkg.Package
}

// Parse decodes the information from a scanner into a package struct.
func (parser *Parser) Parse(scnr scanner.Scanner, result *pkg.Package) (err error) {
	parser.scnr = scnr
	parser.result = result

	for {
		token := parser.scnr.Peak()

		switch {
		case token.IsComment():
			parser.scnr.NextLine()
			continue

		case token.IsClass():
			err = parser.ParseClass()
			if err != nil {
				if err.Error() == "end of scanner source" {
					return nil
				}
				return err
			}
		}

		err = parser.scnr.Next()
		if err != nil {
			if err.Error() == "end of scanner source" {
				return nil
			}
			return err
		}
	}
}
