package parser

import (
	"errors"
)

// ParseFunction returns a string of the build instruction function.
func (p *Parser) ParseFunction() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsFunction() {
		return result, errors.New("called ParseFunction without the beginning token being a function declaration")
	}

	indentLevel := p.scnr.GetIndentLevel()

	for {
		result += p.scnr.PeakLine() + "\n"

		err = p.scnr.NextLine()
		if err != nil {
			return result, err
		}

		if p.scnr.GetIndentLevel() <= indentLevel {
			result += "\n"
			break
		}
	}

	return result, nil
}
