package pkg

import (
	"github.com/DataDrake/cuppa/version"
)

// Package is the internal struct representation of the Spack Package Spec.
type Package struct {
	Name              string
	Description       string
	BlockComment      string
	Homepage          string
	URL               string
	PackageType       string
	Dependencies      []string
	LatestVersion     version.Version
	Versions          []version.Version
	BuildInstructions string
}
