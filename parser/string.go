package parser

import (
	"errors"
	"strings"
)

// ParseString parses and returns the full contents of a string.
func (p *Parser) ParseString() (result string, err error) {
	var prefix string
	token := p.scnr.Peak()
	if !token.IsString() {
		return result, errors.New("called ParseString without the beginning token being a string start")
	}

	token.Data, prefix = trimStringPrefix(token.Data)

	for {
		if token.IsString() {
			trimmed, suffix := trimStringSuffix(token.Data)
			if strings.HasPrefix(suffix, prefix) {
				result += trimmed
				return result, err
			}
			result += token.Data
		}

		result += token.Data + " "

		token, err = p.scnr.Next()
		if err != nil {
			return result, err
		}
	}
}

func trimStringPrefix(token string) (result, prefix string) {
	result = strings.TrimLeftFunc(token, func(input rune) bool {
		return input == '"' || input == '`' || input == '\'' || input == '(' || input == '['
	})
	return result, token[:(len(token) - len(result))]
}

func trimStringSuffix(token string) (result, suffix string) {
	result = strings.TrimRightFunc(token, func(input rune) bool {
		return input == '"' || input == '`' || input == '\'' || input == ',' || input == ')' || input == ']'
	})
	return result, token[len(result):]
}
