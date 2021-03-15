package parser

import (
	"errors"
	"strings"

	"github.com/alecbcs/cuppa/version"
	"github.com/autamus/go-parspack/pkg"
)

// ParseVersion returns the value of a version tuple.
func (p *Parser) ParseVersion() (result pkg.Version, err error) {
	token := p.scnr.Peak()
	if !token.IsVersion() {
		return result, errors.New("called ParseVersion without the beginning token being a version declaration")
	}

	// Parse Version Value
	noprefix := strings.TrimPrefix(strings.ToLower(token.Data), "version('")
	value := strings.TrimSuffix(noprefix, "',")
	result.Value = version.NewVersion(value)
	if strings.Join(result.Value, "") == "N/A" {
		result.Tag = value
	}

	end := false
	for !end {
		token, err = p.scnr.Next()
		if err != nil {
			return result, err
		}

		if strings.HasSuffix(token.Data, ")") {
			end = true
			token.Data = strings.TrimSuffix(token.Data, ")")
		} else {
			token.Data = strings.TrimSuffix(token.Data, ",")
		}
		p.scnr.SetToken(token.Data)

		switch {
		case token.IsChecksum():
			// Parse Checksum
			result.Checksum = token.Data
			break
		case token.IsURL():
			result.URL, err = p.ParseURL()
			if err != nil {
				return result, err
			}
			break
		case token.IsBranch():
			result.Branch, err = p.ParseBranch()
			if err != nil {
				return result, err
			}
			break
		case token.IsSubmodule():
			result.Submodules, err = p.ParseSubmodule()
			if err != nil {
				return result, err
			}
			break
		case token.IsExpand():
			result.Expand, err = p.ParseExpand()
			if err != nil {
				return result, err
			}
			break
		}
	}

	return result, nil
}
