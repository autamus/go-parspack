package parser

import (
	"errors"
	"strings"
)

// ParseClass handles the specific parsing of a package class.
func (p *Parser) ParseClass() (err error) {
	// Set one time switches
	URLSet := false
	GitURLSet := false
	HomepageSet := false

	token := p.scnr.Peak()
	if !token.IsClass() {
		return errors.New("called ParseClass without the beginning token being a class defintion")
	}

	_, err = p.scnr.Next()
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
		// Skip line when comment seen.
		case strings.HasPrefix(token.Data, "#"):
			err = p.scnr.NextLine()
			if err != nil {
				break
			}
			continue

		case strings.HasPrefix(token.Data, `"""`):
			p.result.Description, err = p.ParseString()

		case token.IsBracket():
			err = p.ParseBracket()
			if err != nil {
				return err
			}
			continue

		case token.IsString():
			_, err = p.ParseString()

		case token.IsHomepage() && !HomepageSet:
			p.result.Homepage, err = p.ParseHomepage()
			HomepageSet = true

		case token.IsURL() && !URLSet:
			p.result.URL, err = p.ParseURL()
			URLSet = true

		case token.IsMaintainers():
			p.result.Maintainers, err = p.ParseMaintainers()

		case token.IsGitURL() && !GitURLSet:
			p.result.GitURL, err = p.ParseGitURL()
			GitURLSet = true

		case token.IsVersion():
			version, err := p.ParseVersion()
			if err != nil {
				return err
			}
			p.result.AddVersion(version)

		case token.IsDependency():
			dependency, err := p.ParseDependency()
			if err != nil {
				return err
			}
			p.result.Dependencies = append(p.result.Dependencies, dependency)

		case token.IsFunction():
			result, err := p.ParseFunction()
			p.result.BuildInstructions += result
			if err != nil {
				return err
			}
			continue
		}

		if err != nil {
			return err
		}

		_, err = p.scnr.Next()
		if err != nil {
			break
		}
	}
	return err
}

// ParseClassName take care of parsing the name of a package.
func (p *Parser) ParseClassName() (err error) {
	token := p.scnr.Peak()
	data := strings.Split(strings.TrimRight(token.Data, "):"), "(")
	if len(data) != 2 {
		return errors.New("could not find package name and type")
	}

	p.result.Name = data[0]
	p.result.PackageType = data[1]
	return nil
}
