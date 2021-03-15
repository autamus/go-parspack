package parser

import (
	"errors"
	"strings"
)

// ParseExpand returns the value of the expand variable.
func (p *Parser) ParseExpand() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsExpand() {
		return result, errors.New("called ParseExpand without the beginning token being a expand declaration")
	}

	if strings.HasPrefix(strings.ToLower(token.Data), "expand=") {
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
	return p.ParseBoolean()
}
