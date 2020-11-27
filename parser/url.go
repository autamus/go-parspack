package parser

import "errors"

// ParseURL returns the value of the url variable.
func (p *Parser) ParseURL() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsURL() {
		return result, errors.New("called ParseURL without the beginning token being a URL declaration")
	}

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
	return p.ParseString()
}
