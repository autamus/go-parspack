package parser

import (
	"errors"
	"strings"
)

// ParseDependency returns a list of strings of dependency packages.
func (p *Parser) ParseDependency() (result string, err error) {
	token := p.scnr.Peak()
	if !token.IsDependency() {
		return result, errors.New("called ParseDependencies without the beginning token being a depends_on declaration")
	}

	// Remember if the dependency was only one token wide.
	hadSuffix := false
	if strings.HasSuffix(token.Data, ")") {
		hadSuffix = true
	}

	// Trim left and right sides of dependency token
	token.Data = strings.TrimRight(strings.TrimLeft(strings.ToLower(token.Data), "deps_on("), ")")
	// Save the modified token back to the parser to use with ParseString
	p.scnr.SetToken(token.Data)
	result, err = p.ParseString()
	if err != nil {
		return result, err
	}

	// Record the end of the depedency name versus version/variant info.
	end := strings.IndexFunc(result, versend)
	if end > 0 {
		result = result[:end]
	}

	for {
		if strings.HasSuffix(token.Data, ")") || hadSuffix {
			break
		}

		token, err = p.scnr.Next()
		if err != nil {
			return result, err
		}
	}

	return result, nil
}

// versend returns true at the end of the name of a dependency
func versend(input rune) bool {
	for _, c := range []rune{'@', '~', '+'} {
		if input == c {
			return true
		}
	}
	return false
}
