package parspack

import (
	"bytes"
	"fmt"
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

	rawData := strings.Split(inputRaw, "version(")
	beginning := strings.TrimRight(strings.Split(rawData[0], "url")[0], " ")
	end := strings.SplitN(rawData[len(rawData)-1], "\n", 2)
	result = beginning + fmt.Sprintf("    url      = \"%s\"\n\n", input.URL) + versions + end[1]

	return result, nil
}
