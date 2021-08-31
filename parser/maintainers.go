package parser

import "errors"

// ParseMaintainers returns the value of the homepage variable.
func (p *Parser) ParseMaintainers() (result []string, err error) {
	token := p.scnr.Peak()
	if !token.IsMaintainers() {
		return result, errors.New("called ParseMaintainers without the beginning token being a maintainer declaration")
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
	if token.IsList() {
		return p.ParseList()
	}
	return result, nil
}
