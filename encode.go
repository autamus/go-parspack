package parspack

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/autamus/go-parspack/pkg"
)

// Encode writes the data from a package struct to a String
// to create a spack build spec.
func Encode(input pkg.Package) (result string, err error) {
	t, err := template.New("spec").Funcs(template.FuncMap{
		"printVersion": func(input pkg.Version) string {
			result := strings.Join(input.Value, ".")
			if result == "N/A" {
				return input.Tag
			}
			return result
		},
	}).Parse(pkg.PkgTemplate)
	if err != nil {
		return result, err
	}

	var buff bytes.Buffer
	err = t.Execute(&buff, input)
	if err != nil {
		return result, err
	}

	result = buff.String()

	return result, nil
}
