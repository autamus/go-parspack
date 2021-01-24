package parspack

import (
	"errors"
	"log"
	"testing"

	"github.com/DataDrake/cuppa/version"

	"github.com/autamus/go-parspack/pkg"
)

func TestDecode(t *testing.T) {
	result, err := Decode(`
	# Copyright 2013-2020 Lawrence Livermore National Security, LLC and other
	# Spack Project Developers. See the top-level COPYRIGHT file for details.
	#
	# SPDX-License-Identifier: (Apache-2.0 OR MIT)

	from spack import *


	class Beast2(Package):
		"""BEAST is a cross-platform program for Bayesian inference using MCMC
		of molecular sequences. It is entirely orientated towards rooted,
		time-measured phylogenies inferred using strict or relaxed molecular
		clock models. It can be used as a method of reconstructing phylogenies
		but is also a framework for testing evolutionary hypotheses without
		conditioning on a single tree topology."""

		homepage = "http://beast2.org/"
		url      = "https://github.com/CompEvol/beast2/releases/download/v2.4.6/BEAST.v2.4.6.Linux.tgz"

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
	`)
	if err != nil {
		log.Fatal(err)
	}

	expected := pkg.Package{
		BlockComment: `	# Copyright 2013-2020 Lawrence Livermore National Security, LLC and other
	# Spack Project Developers. See the top-level COPYRIGHT file for details.
	#
	# SPDX-License-Identifier: (Apache-2.0 OR MIT)
`,
		Name:        "Beast2",
		PackageType: "Package",
		Description: `BEAST is a cross-platform program for Bayesian inference using MCMC of molecular sequences. It is entirely orientated towards rooted, time-measured phylogenies inferred using strict or relaxed molecular clock models. It can be used as a method of reconstructing phylogenies but is also a framework for testing evolutionary hypotheses without conditioning on a single tree topology.`,
		Homepage:    "http://beast2.org/",
		URL:         "https://github.com/CompEvol/beast2/releases/download/v2.4.6/BEAST.v2.4.6.Linux.tgz",
		Versions: []pkg.Version{pkg.Version{Value: version.NewVersion("2.5.2"), Checksum: "sha256='2feb2281b4f7cf8f7de1a62de50f52a8678ed0767fc72f2322e77dde9b8cd45f'"},
			pkg.Version{Value: version.NewVersion("2.4.6"), Checksum: "sha256='84029c5680cc22f95bef644824130090f5f12d3d7f48d45cb4efc8e1d6b75e93'", URL: "https://github.com/CompEvol/beast2/releases/download/v2.4.6/BEAST.v2.4.6.Linux.tgz"}},
		LatestVersion: pkg.Version{Value: version.NewVersion("2.5.2"), Checksum: "sha256='2feb2281b4f7cf8f7de1a62de50f52a8678ed0767fc72f2322e77dde9b8cd45f'"},
		Dependencies:  []string{"java"},
		BuildInstructions: `		def setup_run_environment(self, env):
			env.set('BEAST', self.prefix)

		def install(self, spec, prefix):
			install_tree('bin', prefix.bin)
			install_tree('examples', join_path(self.prefix, 'examples'))
			install_tree('images', join_path(self.prefix, 'images'))
			install_tree('lib', prefix.lib)
			install_tree('templates', join_path(self.prefix, 'templates'))
`,
	}

	if result.BlockComment != expected.BlockComment {
		t.Log(result.BlockComment)
		t.Log(expected.BlockComment)
		t.Error(errors.New("result package block comment doesn't match expected"))
	}

	if result.Name != expected.Name {
		t.Error(errors.New("result package name doesn't match expected"))
	}
	if result.PackageType != expected.PackageType {
		t.Error(errors.New("result package type doesn't match expected"))
	}
	if result.Description != expected.Description {
		t.Error(errors.New("result package description doesn't match expected"))
	}
	if result.Homepage != expected.Homepage {
		t.Error(errors.New("result package homepage doesn't match expected"))
	}
	if result.URL != expected.URL {
		t.Error(errors.New("result package URL doesn't match expected"))
	}
	if len(result.Versions) != len(expected.Versions) {
		t.Error(errors.New("result package versions doesn't match expected"))
	} else {
		for i := range expected.Versions {
			if result.Versions[i].Compare(expected.Versions[i]) != 0 {
				t.Error(errors.New("result package versions values don't match expected"))
				t.Log(result.Versions[i])
				t.Log(expected.Versions[i])
			}
			if result.Versions[i].Checksum != expected.Versions[i].Checksum {
				t.Error(errors.New("result package versions checksums don't match expected"))
				t.Log(result.Versions[i])
				t.Log(expected.Versions[i])
			}
			if result.Versions[i].URL != expected.Versions[i].URL {
				t.Error(errors.New("result package versions urls don't match expected"))
				t.Log(result.Versions[i])
				t.Log(expected.Versions[i])
			}
		}
	}
	if result.LatestVersion.Compare(expected.LatestVersion) != 0 {
		t.Error(errors.New("result package LatestVersion doesn't match expected"))
	}
	if len(result.Dependencies) != len(expected.Dependencies) {
		t.Error(errors.New("result package Dependencies don't match expected"))
	} else {
		for i := range expected.Dependencies {
			if result.Dependencies[i] != expected.Dependencies[i] {
				t.Error(errors.New("result package Dependencies don't match expected"))
			}
		}
	}
	if result.BuildInstructions != expected.BuildInstructions {
		t.Log(result.BuildInstructions)
		t.Log(expected.BuildInstructions)
		t.Error(errors.New("result package build instructions don't match expected"))
	}
}
