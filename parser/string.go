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

	token.Data = trimStringPrefix(token.Data)

	for {
		if token.IsString() {
			result += trimStringSuffix(token.Data)
			break
		}

		result += token.Data + " "

		token, err = p.scnr.Next()
		if err != nil {
			return result, err
		}
	}
	return result, err
}

func trimStringPrefix(token string) string {
	return strings.TrimLeftFunc(token, func(input rune) bool {
		return input == '"' || input == '`' || input == '\'' || input == '('
	})
}

func trimStringSuffix(token string) string {
	return strings.TrimRightFunc(token, func(input rune) bool {
		return input == '"' || input == '`' || input == '\'' || input == ',' || input == ')'
	})
}

func trimString(token string) string {
	return strings.TrimFunc(token, func(input rune) bool {
		return input == '"' || input == '`' || input == '\''
	})
}
