package pkg

import "strings"

// Package is the internal struct representation of the Spack Package Spec.
type Package struct {
	Name              string
	Description       string
	BlockComment      string
	Homepage          string
	URL               string
	GitURL            string
	PackageType       string
	Dependencies      []string
	LatestVersion     Version
	Versions          []Version
	BuildInstructions string
}

func (p *Package) containsVersion(input Version) bool {
	if strings.Join(input.Value, "") == "N/A" {
		for _, a := range p.Versions {
			if a.Tag == input.Tag {
				return true
			}
		}
	} else {
		for _, a := range p.Versions {
			if a.Compare(input) == 0 {
				return true
			}
		}
	}
	return false
}
