package parser

import (
	"errors"
	"strings"
)

// ParseBoolean parses and returns the value of a boolean.
func (p *Parser) ParseBoolean() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsBoolean() {
		return result, errors.New("called ParseBoolean without the beginning token being a boolean start")
	}

	token.Data = strings.TrimPrefix(token.Data, "(")
	token.Data = strings.TrimSuffix(token.Data, ")")
	token.Data = strings.TrimSuffix(token.Data, ",")

	return token.Data, err
}
