package pkg

import "github.com/DataDrake/cuppa/version"

// Version is the internal struct representation of the Spack Package Version
type Version struct {
	Value    version.Version
	Checksum string
}

// Compare checks a version against another and returns
// if the other version is smaller than (-1), equal to (0)
// or greater than (1) the current version.
func (v *Version) Compare(other Version) int {
	return v.Value.Compare(other.Value)
}
