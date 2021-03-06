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

	urlExp := regexp.MustCompile(`url\s*=.*`)
	versionExp := regexp.MustCompile(`\s+version\([^\)]*\)`)

	rawData := versionExp.Split(inputRaw, -1)
	rawUrl := urlExp.Split(rawData[0], 2)
	beginning := strings.TrimRight(rawUrl[0], " ")
	end := strings.SplitN(rawData[len(rawData)-1], "\n", 2)
	urlSuffix := ""
	if len(rawUrl) > 1 {
		urlSuffix = strings.TrimPrefix(strings.TrimSuffix(strings.TrimSuffix(rawUrl[1], "    "), "\n"), "\n")
	}
	result = strings.TrimSuffix(beginning, "\n") + "\n"
	if len(input.URL) > 0 {
		result += fmt.Sprintf("    url      = \"%s\"\n", input.URL)
	}
	result += urlSuffix + "\n" + versions + end[1]

	return result, nil
}
