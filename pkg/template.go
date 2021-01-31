package pkg

// PkgTemplate is the common template to generate encoded spack
// package specs.
var PkgTemplate = "" +
	`{{.BlockComment}}

from spack import *


class {{.Name}}({{.PackageType}}):
    {{if .Description}}"""{{.Description}}"""{{end}}

    {{if .Homepage}}homepage = "{{.Homepage}}"{{end}}
    {{if .URL}}url      = "{{.URL}}"{{end}}

` + VersionTemplate + `
{{range $_, $entry := .Dependencies}}    depends_on('{{$entry}}')
{{end}}
{{.BuildInstructions}}
`

// VersionTemplate is the defining template for how versions are
// written to generate an encoded spack package.
var VersionTemplate = "" +
	`{{range $_, $entry := .Versions}}    version('{{printVersion $entry}}', {{$entry.Checksum}}{{if $entry.URL}}, url='{{$entry.URL}}'{{end}})
{{end}}`
