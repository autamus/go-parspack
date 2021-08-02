package parser

import "errors"

// ParseHomepage returns the value of the homepage variable.
func (p *Parser) ParseHomepage() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsHomepage() {
		return result, errors.New("called ParseHomepage without the beginning token being a homepage declaration")
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
	if token.IsString() {
		return p.ParseString()
	}
	return result, nil
}
