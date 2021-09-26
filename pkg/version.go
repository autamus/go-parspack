package pkg

import (
	"regexp"
	"sort"
	"strings"

	"github.com/DataDrake/cuppa/version"
)

// Version is the internal struct representation of the Spack Package Version
type Version struct {
	Value      version.Version
	Checksum   string
	Commit     string
	URL        string
	Tag        string
	Branch     string
	Submodules string
	Expand     string
	Extension  string
}

var IsAlphabetic = regexp.MustCompile(`^[a-zA-Z/]+$`).MatchString

// AddVersion appends a new version to the package struct if it doesn't already
// exist and sets the latest version to the input version if it is now the latest.
func (p *Package) AddVersion(input Version) {
	if !p.containsVersion(input) {
		// Check filetype ending of URL to determine if expand = True/False
		url := input.URL
		if url == "" {
			url = p.URL
		}
		if input.Expand == "" && input.URL != "" &&
			!strings.HasSuffix(url, ".tar.gz") &&
			!strings.HasSuffix(url, ".tar.bz2") &&
			!strings.HasSuffix(url, ".tgz") &&
			!strings.HasSuffix(url, ".zip") {
			input.Expand = "False"
		}
		p.Versions = append(p.Versions, input)
		if !IsAlphabetic(input.Value.String()) &&
			(p.LatestVersion.Value == nil || p.LatestVersion.Compare(input) > 0) {
			p.LatestVersion = input
		}

		// Sort versions from high to low.
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
