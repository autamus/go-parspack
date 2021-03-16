package parspack

import (
	"testing"
)

func TestPatch(t *testing.T) {
	spackPackage := `# Copyright 2013-2021 Lawrence Livermore National Security, LLC and other
# Spack Project Developers. See the top-level COPYRIGHT file for details.
#
# SPDX-License-Identifier: (Apache-2.0 OR MIT)

from spack import *
import glob
import os.path
import re

class Picard(Package):
    """Picard is a set of command line tools for manipulating high-throughput
        sequencing (HTS) data and formats such as SAM/BAM/CRAM and VCF.
    """

    homepage = "http://broadinstitute.github.io/picard/"
    url      = "https://github.com/broadinstitute/picard/releases/download/2.9.2/picard.jar"
    _urlfmt  = "https://github.com/broadinstitute/picard/releases/download/{0}/picard.jar"
    _oldurlfmt = 'https://github.com/broadinstitute/picard/releases/download/{0}/picard-tools-{0}.zip'

    # They started distributing a single jar file at v2.6.0, prior to
    # that it was a .zip file with multiple .jar and .so files
    version('2.24.0', sha256='70e91039bccc6f6db60f18c41713218a8cdf45f591f02c1012c062152b27cd7b', expand=False)
    version('2.20.8', sha256='aff92d618ee9e6bafc1ab4fbfa89fc557a0dbe596ae4b92fe3bf93ebf95c7105', expand=False)
    version('2.19.0', sha256='f97fc3f7a73b55cceea8b6a6488efcf1b2fbf8cad61d88645704ddd45a8c5950', expand=False)
    version('2.18.3', sha256='0e0fc45d9a822ee9fa562b3bb8f1525a439e4cd541429a1a4ca4646f37189b70', expand=False)
    version('2.18.0', sha256='c4c64b39ab47968e4afcaf1a30223445ee5082eab31a03eee240e54c2e9e1dce', expand=False)
    version('2.17.0', sha256='ffea8bf90e850919c0833e9dcc16847d40090a1ef733c583a710a3282b925998', expand=False)
    version('2.16.0', sha256='01cf3c930d2b4841960497491512d327bf669c1ed2e753152e1e651a27288c2d', expand=False)
    version('2.15.0', sha256='dc3ff74c704954a10796b390738078617bb0b0fef15731e9d268ed3b26c6a285', expand=False)
    version('2.13.2', sha256='db7749f649e8423053fb971e6af5fb8a9a1a797cb1f20fef1100edf9f82f6f94', expand=False)
    version('2.10.0', sha256='e256d5e43656b7d8be454201a7056dce543fe9cbeb30329a0d8c22d28e655775', expand=False)
    version('2.9.4', sha256='0ecee9272bd289d902cc18053010f0364d1696e7518ac92419a99b2f0a1cf689', expand=False)
    version('2.9.3', sha256='6cca26ce5094b25a82e1a8d646920d584c6db5f009476196dc285be6522e00ce', expand=False)
    version('2.9.2', sha256='05714b9743a7685a43c94a93f5d03aa4070d3ab6e20151801301916d3e546eb7', expand=False)
    version('2.9.0', sha256='9a57f6bd9086ea0f5f1a6d9d819459854cb883bb8093795c916538ed9dd5de64', expand=False)
    version('2.8.3', sha256='97a4b6c8927c8cb5f3450630c9b39bf210ced8c271f198119118ce1c24b8b0f6', expand=False)
    version('2.6.0', sha256='671d9e86e6bf0c28ee007aea55d07e2456ae3a57016491b50aab0fd2fd0e493b', expand=False)
    version('1.140', sha256='0d27287217413db6b846284c617d502eaa578662dcb054a7017083eab9c54438')

    depends_on('java@8:', type='run')

    def install(self, spec, prefix):
        mkdirp(prefix.bin)
        # The list of files to install varies with release...
        # ... but skip the spack-build-{env}out.txt files.
        files = [x for x in glob.glob("*") if not re.match("^spack-", x)]
        for f in files:
            install(f, prefix.bin)

        # Set up a helper script to call java on the jar file,
        # explicitly codes the path for java and the jar file.
        script_sh = join_path(os.path.dirname(__file__), "picard.sh")
        script = prefix.bin.picard
        install(script_sh, script)
        set_executable(script)

        # Munge the helper script to explicitly point to java and the
        # jar file.
        java = self.spec['java'].prefix.bin.java
        kwargs = {'ignore_absent': False, 'backup': False, 'string': False}
        filter_file('^java', java, script, **kwargs)
        filter_file('picard.jar', join_path(prefix.bin, 'picard.jar'),
                    script, **kwargs)

    def setup_run_environment(self, env):
        """The Picard docs suggest setting this as a convenience."""
        env.prepend_path('PICARD', join_path(self.prefix.bin, 'picard.jar'))

    def url_for_version(self, version):
        if version < Version('2.6.0'):
            return self._oldurlfmt.format(version)
        else:
            return self._urlfmt.format(version)

	`
	result, err := Decode(spackPackage)
	if err != nil {
		t.Fatal(err)
	}

	output, err := PatchVersion(result, spackPackage)
	if output != spackPackage {
		t.Error(output)
		t.Error(spackPackage)
	}
}

func TestPatchComplicatedUrl(t *testing.T) {
	spackPackage := `# Copyright 2013-2021 Lawrence Livermore National Security, LLC and other
# Spack Project Developers. See the top-level COPYRIGHT file for details.
#
# SPDX-License-Identifier: (Apache-2.0 OR MIT)

from spack import *


class Htslib(AutotoolsPackage):
    """C library for high-throughput sequencing data formats."""

    homepage = "https://github.com/samtools/htslib"
    url      = "https://github.com/samtools/htslib/releases/download/1.10.2/htslib-1.10.2.tar.bz2"

    version('1.11', sha256='cffadd9baa6fce27b8fe0b01a462b489f06a5433dfe92121f667f40f632538d7')
    version('1.10.2', sha256='e3b543de2f71723830a1e0472cf5489ec27d0fbeb46b1103e14a11b7177d1939')
    version('1.9', sha256='e04b877057e8b3b8425d957f057b42f0e8509173621d3eccaedd0da607d9929a')
    version('1.8', sha256='c0ef1eec954a98cc708e9f99f6037db85db45670b52b6ab37abcc89b6c057ca1')
    version('1.7', sha256='be3d4e25c256acdd41bebb8a7ad55e89bb18e2fc7fc336124b1e2c82ae8886c6')
    version('1.6', sha256='9588be8be0c2390a87b7952d644e7a88bead2991b3468371347965f2e0504ccb')
    version('1.5', sha256='a02b515ea51d86352b089c63d778fb5e8b9d784937cf157e587189cb97ad922d')
    version('1.4', sha256='5cfc8818ff45cd6e924c32fec2489cb28853af8867a7ee8e755c4187f5883350')
    version('1.3.1', sha256='49d53a2395b8cef7d1d11270a09de888df8ba06f70fe68282e8235ee04124ae6')
    version('1.2', sha256='125c01421d5131afb4c3fd2bc9c7da6f4f1cd9ab5fc285c076080b9aca24bffc')

    variant('libcurl',
            default = True,
            description = 'Enable libcurl-based support for http/https/etc URLs,'
            ' for versions >= 1.3. This also enables S3 and GCS support.')

    depends_on('zlib')
    depends_on('bzip2', when='@1.4:')
    depends_on('xz', when='@1.4:')
    depends_on('curl', when='@1.3:+libcurl')

    depends_on('m4', when="@1.2")
    depends_on('autoconf', when="@1.2")
    depends_on('automake', when="@1.2")
    depends_on('libtool', when="@1.2")

    # v1.2 uses the automagically assembled tarball from .../archive/...
    # everything else uses the tarballs uploaded to the release
    def url_for_version(self, version):
        if version.string == '1.2':
            return 'https://github.com/samtools/htslib/archive/1.2.tar.gz'
        else:
            url = "https://github.com/samtools/htslib/releases/download/{0}/htslib-{0}.tar.bz2"
            return url.format(version.dotted)

    def configure_args(self):
        spec = self.spec
        args = []

        if spec.satisfies('@1.3:'):
            args.extend(self.enable_or_disable('libcurl'))

        return args
    
	`
	result, err := Decode(spackPackage)
	if err != nil {
		t.Fatal(err)
	}

	output, err := PatchVersion(result, spackPackage)
	if output != spackPackage {
		t.Error(output)
		t.Error(spackPackage)
	}
}
