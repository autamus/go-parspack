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
