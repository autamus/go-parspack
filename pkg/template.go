package pkg

// PkgTemplate is the common template to generate encoded spack
// package specs.
var PkgTemplate = "" +
	`{{.BlockComment}}

from spack import *


class {{.Name}}({{.PackageType}}):
	"""{{.Description}}"""

	homepage = "{{.Homepage}}"
	{{if .URL}}url      = "{{.URL}}"{{end}}

{{range $_, $entry := .Versions}}	version('{{printVersion $entry}}', {{$entry.Checksum}}{{if $entry.URL}}, url='{{$entry.URL}}'{{end}})
{{end}}
	{{range $_, $entry := .Dependencies}}depends_on('{{$entry}}'){{end}}

{{.BuildInstructions}}
`
