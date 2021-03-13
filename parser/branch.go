package parser

import (
	"errors"
	"strings"
)

// ParseBranch returns the value of the branch variable.
func (p *Parser) ParseBranch() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsBranch() {
		return result, errors.New("called ParseBranch without the beginning token being a branch declaration")
	}

	if strings.HasPrefix(strings.ToLower(token.Data), "branch=") {
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
