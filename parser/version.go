package parser

import (
	"errors"
	"strings"

	"github.com/DataDrake/cuppa/version"
	"github.com/autamus/go-parspack/pkg"
)

// ParseVersion returns the value of a version tuple.
func (p *Parser) ParseVersion() (result pkg.Version, err error) {
	// Watch for the end of the version.
	end := false
	token := p.scnr.Peak()
	if !token.IsVersion() {
		return result, errors.New("called ParseVersion without the beginning token being a version declaration")
	}

	// Parse Version Value
	noprefix := strings.TrimPrefix(strings.ToLower(token.Data), "version(")
	if noprefix != "" {
		p.scnr.SetToken(noprefix)
	} else {
		p.scnr.Next()
	}

	value, err := p.ParseString()
	if err != nil {
		return result, err
	}

	if strings.HasSuffix(noprefix, ")") {
		end = true
	}
	result.Value = version.NewVersion(value)

	// Check for N/A version value
	if result.Value.String() == "N/A" {
		result.Value = []string{value}
	}

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
		case token.IsURL():
			result.URL, err = p.ParseURL()
			if err != nil {
				return result, err
			}
		case token.IsBranch():
			result.Branch, err = p.ParseBranch()
			if err != nil {
				return result, err
			}
		case token.IsSubmodule():
			result.Submodules, err = p.ParseSubmodule()
			if err != nil {
				return result, err
			}
		case token.IsExpand():
			result.Expand, err = p.ParseExpand()
			if err != nil {
				return result, err
			}
		case token.IsCommit():
			result.Commit, err = p.ParseCommit()
			if err != nil {
				return result, err
			}
		}
	}

	return result, nil
}
