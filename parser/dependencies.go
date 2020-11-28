package parser

import (
	"errors"
	"strings"
)

// ParseDependencies returns a list of strings of dependency packages.
func (p *Parser) ParseDependencies() (result []string, err error) {
	token := p.scnr.Peak()
	if !token.IsDependency() {
		return result, errors.New("called ParseDependencies without the beginning token being a depends_on declaration")
	}

	token.Data = strings.TrimLeft(strings.ToLower(token.Data), "depends_on(")

	for {
		if strings.HasSuffix(token.Data, ")") {
			data := strings.TrimRight(token.Data, ")")
			result = append(result, trimString(data))
			break
		}

		data := strings.TrimRight(token.Data, ",")
		result = append(result, trimString(data))

		token, err = p.scnr.Next()
		if err != nil {
			return result, err
		}
	}

	return result, nil
}
