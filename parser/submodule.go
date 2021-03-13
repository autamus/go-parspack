package parser

import (
	"errors"
	"strings"
)

// ParseSubmodule returns the value of the branch variable.
func (p *Parser) ParseSubmodule() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsSubmodule() {
		return result, errors.New("called ParseSubmodule without the beginning token being a branch declaration")
	}

	if strings.HasPrefix(strings.ToLower(token.Data), "submodules=") {
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
