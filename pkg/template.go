package pkg

// PkgTemplate is the common template to generate encoded spack
// package specs.
var PkgTemplate = "" +
	`{{.BlockComment}}
	
from spack import *


class {{.Name}}({{.PackageType}}):
	"""{{.Description}}"""

	homepage = "{{.Homepage}}"
	url      = "{{.URL}}"

	{{range $_, $entry := .Versions}}version('{{printVersion $entry}}', {{$entry.Checksum}})
	{{end}}
	{{range $_, $entry := .Dependencies}}depends_on('{{$entry}}'){{end}}

{{.BuildInstructions}}
	`
