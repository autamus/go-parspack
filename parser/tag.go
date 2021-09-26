package parser

import (
	"errors"
	"strings"
)

// ParseTag returns the value of the branch variable.
func (p *Parser) ParseTag() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsTag() {
		return result, errors.New("called ParseTag without the beginning token being a tag declaration")
	}

	if strings.HasPrefix(strings.ToLower(token.Data), "tag=") {
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
