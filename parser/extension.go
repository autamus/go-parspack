package parser

import (
	"errors"
	"strings"
)

// ParseExtension returns the value of the branch variable.
func (p *Parser) ParseExtension() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsExtension() {
		return result, errors.New("called ParseExtension without the beginning token being a extension declaration")
	}

	if strings.HasPrefix(strings.ToLower(token.Data), "extension=") {
		token.Data = strings.SplitN(token.Data, "=", 2)[1]
		p.scnr.SetToken(token.Data)
	} else {
		prev := token
		token, err = p.scnr.Next()
		if err != nil {
			return result, err
		}
		if token.Data != "=" {
			return result, errors.New("expecting '=' after" + prev.Data)
		}

		token, err = p.scnr.Next()
		if err != nil {
			return result, err
		}
	}
	return p.ParseString()
}
