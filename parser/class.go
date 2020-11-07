package parser

import (
	"errors"
	"strings"
)

// ParseClass handles the specific parsing of a package class.
func (p *Parser) ParseClass() (err error) {
	token := p.scnr.Peak()
	if !token.IsClass() {
		return errors.New("called ParseClass without the beginning token being a class defintion")
	}

	err = p.scnr.Next()
	if err != nil {
		return err
	}
	err = p.ParseClassName()
	if err != nil {
		return err
	}

	err = p.scnr.NextLine()
	if err != nil {
		return err
	}

	for {
		token := p.scnr.Peak()

		switch {
		case token.IsString():
			p.result.Description, err = p.ParseString()
			if err != nil {
				return err
			}

		}
	}

	return nil
}

// ParseClassName take care of parsing the name of a package.
func (p *Parser) ParseClassName() (err error) {
	token := p.scnr.Peak()
	data := strings.Split(strings.TrimRight(token.Data, ")"), "(")
	if len(data) != 2 {
		return errors.New("could not find package name and type")
	}

	p.result.Name = data[0]
	p.result.PackageType = data[1]
	return nil
}
