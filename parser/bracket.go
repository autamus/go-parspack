package parser

import (
	"errors"
	"strings"
)

// ParseString parses and returns the full contents of a string.
func (p *Parser) ParseBracket() (err error) {
	token := p.scnr.Peak()
	if !token.IsBracket() {
		return errors.New("called ParseString without the beginning token being a bracket start")
	}

	token.Data = trimBracketPrefix(token.Data)

	if token.IsBracket() {
		token.Data = trimBracketSuffix(token.Data)
	}

	p.scnr.SetToken(token.Data)
	return nil
}

func trimBracketPrefix(token string) string {
	return strings.TrimLeftFunc(token, func(input rune) bool {
		return input == '['
	})
}

func trimBracketSuffix(token string) string {
	return strings.TrimRightFunc(token, func(input rune) bool {
		return input == ']'
	})
}
