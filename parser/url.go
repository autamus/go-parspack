package parser

import (
	"errors"
	"strings"
)

// ParseURL returns the value of the url variable.
func (p *Parser) ParseURL() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsURL() {
		return result, errors.New("called ParseURL without the beginning token being a URL declaration")
	}

	if strings.HasPrefix(strings.ToLower(token.Data), "url=") {
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

// ParseGitURL returns the value of the url variable.
func (p *Parser) ParseGitURL() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsGitURL() {
		return result, errors.New("called ParseURL without the beginning token being a URL declaration")
	}

	if strings.HasPrefix(strings.ToLower(token.Data), "git=") {
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
