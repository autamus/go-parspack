package parser

import (
	"errors"
	"strings"
)

// ParseList parses and returns the full contents of a string.
func (p *Parser) ParseList() (result []string, err error) {
	token := p.scnr.Peak()
	if !token.IsList() {
		return result, errors.New("called ParseList without the beginning token being a list start")
	}

	token.Data = trimListPrefix(token.Data)

	for {
		if token.IsList() {
			trimmed := trimListSuffix(token.Data)
			p.scnr.SetToken(trimmed)

			token = p.scnr.Peak()
			if token.IsString() {
				str, err := p.ParseString()
				if err != nil {
					return result, err
				}
				result = append(result, str)
			}
			return result, err
		}

		token = p.scnr.Peak()
		if token.IsString() {
			str, err := p.ParseString()
			if err != nil {
				return result, err
			}
			result = append(result, str)
		}

		token, err = p.scnr.Next()
		if err != nil {
			return result, err
		}
	}
}

func trimListPrefix(token string) (result string) {
	result = strings.TrimLeftFunc(token, func(input rune) bool {
		return input == '['
	})
	return result
}

func trimListSuffix(token string) (result string) {
	result = strings.TrimRightFunc(token, func(input rune) bool {
		return input == ']'
	})
	return result
}
