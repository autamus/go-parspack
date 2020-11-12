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
func (p *Parser) Parse(scnr scanner.Scanner, result *pkg.Package) (err error) {
	p.scnr = scnr
	p.result = result

	// Parse Beginning Block Comment.
	for {
		token := p.scnr.Peak()
		if !token.IsComment() {
			break
		}
		p.result.BlockComment = p.result.BlockComment + p.scnr.PeakLine() + "\n"
		p.scnr.NextLine()
	}

	for {
		token := p.scnr.Peak()

		switch {
		case token.IsComment():
			p.scnr.NextLine()
			continue

		case token.IsClass():
			err = p.ParseClass()
			if err != nil {
				if err.Error() == "end of scanner source" {
					return nil
				}
				return err
			}
		}

		_, err = p.scnr.Next()
		if err != nil {
			if err.Error() == "end of scanner source" {
				return nil
			}
			return err
		}
	}
}
