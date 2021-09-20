package parspack

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	"github.com/autamus/go-parspack/pkg"
)

// PatchVersion updates the version and latest URL of a file with the
// information from a Package struct without re-encoding the entire
// file.
func PatchVersion(input pkg.Package, inputRaw string) (result string, err error) {
	t, err := template.New("spec").Funcs(template.FuncMap{
		"printVersion": func(input pkg.Version) string {
			return strings.Join(input.Value, ".")
		},
	}).Parse(pkg.VersionTemplate)
	if err != nil {
		return result, err
	}

	var buff bytes.Buffer
	err = t.Execute(&buff, input)
	if err != nil {
		return result, err
	}

	versions := buff.String()
	urlExp := regexp.MustCompile(`\n    url\s*=.*".*"\n`)
	versionExp := regexp.MustCompile(`\n(    version\([^\)]*\)(\n|.)*)*(    version\([^\)]*\))+`)

	if len(input.Versions) > 0 {
		inputRaw = versionExp.ReplaceAllString(inputRaw, "\n"+versions)
	}

	if input.URL != "" {
		inputRaw = urlExp.ReplaceAllString(inputRaw, fmt.Sprintf("\n    url      = \"%s\"\n", input.URL))
	}

	return inputRaw, nil
}
