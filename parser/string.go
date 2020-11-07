package parser

import (
	"errors"
)

// ParseString parses and returns the full contents of a string.
func (p *Parser) ParseString() (result string, err error) {
	// var data []string

	token := p.scnr.Peak()
	if !token.IsString() {
		return result, errors.New("called ParseString without the beginning token being a string start")
	}
	return result, err
}
