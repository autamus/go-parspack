package pkg

// Package is the internal struct representation of the Spack Package Spec.
type Package struct {
	Name              string
	Description       string
	BlockComment      string
	Homepage          string
	URL               string
	PackageType       string
	Dependencies      []string
	LatestVersion     Version
	Versions          []Version
	BuildInstructions string
}

// AddVersion appends a new version to the package struct if it doesn't already
// exist and sets the latest version to the input version if it is now the latest.
func (p *Package) AddVersion(input Version) {
	if !p.contains(input) {
		p.Versions = append(p.Versions, input)
		if p.LatestVersion.Value == nil || p.LatestVersion.Compare(input) < 0 {
			p.LatestVersion = input
		}
	}
}

func (p *Package) contains(input Version) bool {
	for _, a := range p.Versions {
		if a.Compare(input) == 0 {
			return true
		}
	}
	return false
}
