package pkg

// PkgTemplate is the common template to generate encoded spack
// package specs.
var PkgTemplate = "" +
	`{{.BlockComment}}

from spack import *


class {{.Name}}({{.PackageType}}):
    {{if .Description}}"""{{.Description}}"""{{end}}

    {{if .Homepage}}homepage = "{{.Homepage}}"{{end}}
    {{if gt (len .URL) 0}}url      = "{{.URL}}"{{end}}

` + VersionTemplate + `
{{range $_, $entry := .Dependencies}}    depends_on('{{$entry}}')
{{end}}
{{.BuildInstructions}}
`

// VersionTemplate is the defining template for how versions are
// written to generate an encoded spack package.
var VersionTemplate = "" +
	`{{range $_, $entry := .Versions}}    version('{{printVersion $entry}}'{{if $entry.Checksum}}, {{$entry.Checksum}}{{end}}{{if $entry.Branch}}, branch='{{$entry.Branch}}'{{end}}{{if $entry.Commit}}, commit='{{$entry.Commit}}'{{end}}{{if $entry.Submodules}}, submodules={{$entry.Submodules}}{{end}}{{if $entry.Expand}}, expand={{$entry.Expand}}{{end}}{{if $entry.URL}}, url='{{$entry.URL}}'{{end}}{{if $entry.Extension}}, extension='{{$entry.Extension}}'{{end}}{{if $entry.Tag}}, tag='{{$entry.Tag}}'{{end}})
{{end}}`
