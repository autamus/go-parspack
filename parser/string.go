package parser

import (
	"errors"
	"strings"
)

// ParseString parses and returns the full contents of a string.
func (p *Parser) ParseString() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsString() {
		return result, errors.New("called ParseString without the beginning token being a string start")
	}

	result = trimString(token.Data)

	for {
		token, err = p.scnr.Next()
		if err != nil {
			return result, err
		}

		if token.IsString() {
			result += " " + trimString(token.Data)
			break
		}

		result += " " + token.Data
	}

	return result, err
}

func trimString(token string) string {
	return strings.TrimFunc(token, func(input rune) bool {
		return input == '"' || input == '`' || input == '\''
	})
}
