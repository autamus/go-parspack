package pkg

import (
	"sort"

	"github.com/DataDrake/cuppa/version"
)

// Version is the internal struct representation of the Spack Package Version
type Version struct {
	Value    version.Version
	Checksum string
	URL      string
}

// AddVersion appends a new version to the package struct if it doesn't already
// exist and sets the latest version to the input version if it is now the latest.
func (p *Package) AddVersion(input Version) {
	if !p.containsVersion(input) {
		p.Versions = append(p.Versions, input)
		if p.LatestVersion.Value == nil || p.LatestVersion.Compare(input) > 0 {
			p.LatestVersion = input
		}

		sort.Slice(p.Versions[:], func(i, j int) bool {
			return p.Versions[i].Compare(p.Versions[j]) < 0
		})
	}
}

// Compare checks a version against another and returns
// if the other version is smaller than (-1), equal to (0)
// or greater than (1) the current version.
func (v *Version) Compare(other Version) int {
	return v.Value.Compare(other.Value)
}
