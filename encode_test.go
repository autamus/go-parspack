package parspack

import (
	"testing"

	"github.com/DataDrake/cuppa/version"
	"github.com/autamus/go-parspack/pkg"
)

func TestEncode(t *testing.T) {
	packg := pkg.Package{
		BlockComment: `# Copyright 2013-2020 Lawrence Livermore National Security, LLC and other
# Spack Project Developers. See the top-level COPYRIGHT file for details.
#
# SPDX-License-Identifier: (Apache-2.0 OR MIT)`,
		Name:        "Beast2",
		PackageType: "Package",
		Description: `BEAST is a cross-platform program for Bayesian inference using MCMC of molecular sequences. It is entirely orientated towards rooted, time-measured phylogenies inferred using strict or relaxed molecular clock models. It can be used as a method of reconstructing phylogenies but is also a framework for testing evolutionary hypotheses without conditioning on a single tree topology.`,
		Homepage:    "http://beast2.org/",
		URL:         "https://github.com/CompEvol/beast2/releases/download/v2.4.6/BEAST.v2.4.6.Linux.tgz",
		Versions: []pkg.Version{{Value: version.NewVersion("master"), Tag: "main", Branch: "main", Submodules: "True"}, {Value: version.NewVersion("2.5.2"), Checksum: "sha256='2feb2281b4f7cf8f7de1a62de50f52a8678ed0767fc72f2322e77dde9b8cd45f'"},
			{Value: version.NewVersion("2.4.6"), Checksum: "sha256='84029c5680cc22f95bef644824130090f5f12d3d7f48d45cb4efc8e1d6b75e93'", URL: "https://github.com/CompEvol/beast2/releases/download/v2.4.6/BEAST.v2.4.6.Linux.tgz"}},
		LatestVersion: pkg.Version{Value: version.NewVersion("2.5.2"), Checksum: "sha256='2feb2281b4f7cf8f7de1a62de50f52a8678ed0767fc72f2322e77dde9b8cd45f'"},
		Dependencies:  []string{"java"},
		BuildInstructions: `    def setup_run_environment(self, env):
        env.set('BEAST', self.prefix)

    def install(self, spec, prefix):
        install_tree('bin', prefix.bin)
        install_tree('examples', join_path(self.prefix, 'examples'))
        install_tree('images', join_path(self.prefix, 'images'))
        install_tree('lib', prefix.lib)
        install_tree('templates', join_path(self.prefix, 'templates'))
`,
	}

	expected := `# Copyright 2013-2020 Lawrence Livermore National Security, LLC and other
# Spack Project Developers. See the top-level COPYRIGHT file for details.
#
# SPDX-License-Identifier: (Apache-2.0 OR MIT)

from spack import *


class Beast2(Package):
    """BEAST is a cross-platform program for Bayesian inference using MCMC of molecular sequences. It is entirely orientated towards rooted, time-measured phylogenies inferred using strict or relaxed molecular clock models. It can be used as a method of reconstructing phylogenies but is also a framework for testing evolutionary hypotheses without conditioning on a single tree topology."""

    homepage = "http://beast2.org/"
    url      = "https://github.com/CompEvol/beast2/releases/download/v2.4.6/BEAST.v2.4.6.Linux.tgz"

    version('main', branch='main', submodules=True)
    version('2.5.2', sha256='2feb2281b4f7cf8f7de1a62de50f52a8678ed0767fc72f2322e77dde9b8cd45f')
    version('2.4.6', sha256='84029c5680cc22f95bef644824130090f5f12d3d7f48d45cb4efc8e1d6b75e93', url='https://github.com/CompEvol/beast2/releases/download/v2.4.6/BEAST.v2.4.6.Linux.tgz')

    depends_on('java')

    def setup_run_environment(self, env):
        env.set('BEAST', self.prefix)

    def install(self, spec, prefix):
        install_tree('bin', prefix.bin)
        install_tree('examples', join_path(self.prefix, 'examples'))
        install_tree('images', join_path(self.prefix, 'images'))
        install_tree('lib', prefix.lib)
        install_tree('templates', join_path(self.prefix, 'templates'))

`

	result, err := Encode(packg)
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Log(result)
		t.Log(expected)
		t.Error("Result and Expected Do Not Match")
	}
}
