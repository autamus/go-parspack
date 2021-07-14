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
    version('develop', commit='xfdss46kjsldkfu8')
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
    version('1.140', sha256='0d27287217413db6b846284c617d502eaa578662dcb054a7017083eab9c54438', expand=True)

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

	output, _ := PatchVersion(result, spackPackage)
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

    version('develop', submodules=True)
    version('master')
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

	output, _ := PatchVersion(result, spackPackage)
	if output != spackPackage {
		t.Error(output)
		t.Error(spackPackage)
	}
}

func TestPatchNoUrl(t *testing.T) {
	spackPackage := `# Copyright 2013-2021 Lawrence Livermore National Security, LLC and other
# Spack Project Developers. See the top-level COPYRIGHT file for details.
#
# SPDX-License-Identifier: (Apache-2.0 OR MIT)
import os.path
import shutil
import sys
import tempfile

import spack.util.environment


class Octave(AutotoolsPackage, GNUMirrorPackage):
    """GNU Octave is a high-level language, primarily intended for numerical
    computations.

    It provides a convenient command line interface for solving linear and
    nonlinear problems numerically, and for performing other numerical
    experiments using a language that is mostly compatible with Matlab.
    It may also be used as a batch-oriented language.
    """

    homepage = "https://www.gnu.org/software/octave/"
    gnu_mirror_path = "octave/octave-4.0.0.tar.gz"
    maintainers = ['mtmiller']

    extendable = True

    version('6.3.0', sha256='232065f3a72fc3013fe9f17f429a3df69d672c1f6b6077029a31c8f3cd58a66e')
    version('6.2.0', sha256='457d1fda8634a839e2fd7cfc55b98bd56f36b6ae73d31bb9df43dde3012caa7c')
    version('6.1.0', sha256='6ff34e401658622c44094ecb67e497672e4337ca2d36c0702d0403ecc60b0a57')
    version('5.2.0', sha256='2fea62b3c78d6f38e9451da8a4d26023840725977dffee5250d3d180f56595e1')
    version('5.1.0', sha256='e36b1124cac27c7caa51cc57de408c31676d5f0096349b4d50b57bfe1bcd7495')
    version('4.4.1', sha256='09fbd0f212f4ef21e53f1d9c41cf30ce3d7f9450fb44911601e21ed64c67ae97')
    version('4.4.0', sha256='72f846379fcec7e813d46adcbacd069d72c4f4d8f6003bcd92c3513aafcd6e96')
    version('4.2.2', sha256='77b84395d8e7728a1ab223058fe5e92dc38c03bc13f7358e6533aab36f76726e')
    version('4.2.1', sha256='80c28f6398576b50faca0e602defb9598d6f7308b0903724442c2a35a605333b')
    version('4.2.0', sha256='443ba73782f3531c94bcf016f2f0362a58e186ddb8269af7dcce973562795567')
    version('4.0.2', sha256='39cd8fd36c218fc00adace28d74a6c7c9c6faab7113a5ba3c4372324c755bdc1')
    version('4.0.0', sha256='4c7ee0957f5dd877e3feb9dfe07ad5f39b311f9373932f0d2a289dc97cca3280')

    # patches
    # see https://savannah.gnu.org/bugs/?50234
    patch('patch_4.2.1_inline.diff', when='@4.2.1')

    # Variants
    variant('readline',   default=True)
    variant('arpack',     default=False)
    variant('curl',       default=False)
    variant('fftw',       default=False)
    variant('fltk',       default=False)
    variant('fontconfig', default=False)
    variant('freetype',   default=False)
    variant('glpk',       default=False)
    variant('gl2ps',      default=False)
    variant('gnuplot',    default=False)
    variant('magick',     default=False)
    variant('hdf5',       default=False)
    variant('jdk',        default=False)
    variant('llvm',       default=False)
    variant('opengl',     default=False)
    variant('qhull',      default=False)
    variant('qrupdate',   default=False)
    variant('qscintilla', default=False)
    variant('qt',         default=False)
    variant('suitesparse', default=False)
    variant('zlib',       default=False)

    # Required dependencies
    depends_on('blas')
    depends_on('lapack')
    # Octave does not configure with sed from darwin:
    depends_on('sed', when=sys.platform == 'darwin', type='build')
    depends_on('pcre')
    depends_on('pkgconfig', type='build')

    # Strongly recommended dependencies
    depends_on('readline',     when='+readline')

    # Optional dependencies
    depends_on('arpack-ng',    when='+arpack')
    depends_on('curl',         when='+curl')
    depends_on('fftw',         when='+fftw')
    depends_on('fltk',         when='+fltk')
    depends_on('fontconfig',   when='+fontconfig')
    depends_on('freetype',     when='+freetype')
    depends_on('glpk',         when='+glpk')
    depends_on('gl2ps',        when='+gl2ps')
    depends_on('gnuplot',      when='+gnuplot')
    depends_on('imagemagick',  when='+magick')
    depends_on('hdf5',         when='+hdf5')
    depends_on('java',         when='+jdk')        # TODO: requires Java 6 ?
    depends_on('llvm',         when='+llvm')
    depends_on('gl',           when='+opengl')
    depends_on('gl',           when='+fltk')
    depends_on('qhull',        when='+qhull')
    depends_on('qrupdate',     when='+qrupdate')
    depends_on('qscintilla',   when='+qscintilla')
    depends_on('qt+opengl',    when='+qt')
    depends_on('suite-sparse', when='+suitesparse')
    depends_on('zlib',         when='+zlib')

    def patch(self):
        # Filter mkoctfile.in.cc to use underlying compilers and not
        # Spack compiler wrappers. We are patching the template file
        # and not mkoctfile.cc since the latter is generated as part
        # of the build.
        mkoctfile_in = os.path.join(
            self.stage.source_path, 'src', 'mkoctfile.in.cc'
        )
        quote = lambda s: '"' + s + '"'
        entries_to_patch = {
            r'%OCTAVE_CONF_MKOCTFILE_CC%': quote(self.compiler.cc),
            r'%OCTAVE_CONF_MKOCTFILE_CXX%': quote(self.compiler.cxx),
            r'%OCTAVE_CONF_MKOCTFILE_F77%': quote(self.compiler.f77),
            r'%OCTAVE_CONF_MKOCTFILE_DL_LD%': quote(self.compiler.cxx),
            r'%OCTAVE_CONF_MKOCTFILE_LD_CXX%': quote(self.compiler.cxx)
        }

        for pattern, subst in entries_to_patch.items():
            filter_file(pattern, subst, mkoctfile_in)

    @run_after('install')
    @on_package_attributes(run_tests=True)
    def check_mkoctfile_works_outside_of_build_env(self):
        # Check that mkoctfile is properly configured and can compile
        # Octave extensions outside of the build env
        mkoctfile = Executable(os.path.join(self.prefix, 'bin', 'mkoctfile'))
        helloworld_cc = os.path.join(
            os.path.dirname(__file__), 'helloworld.cc'
        )
        tmp_dir = tempfile.mkdtemp()
        shutil.copy(helloworld_cc, tmp_dir)

        # We need to unset these variables since we are still within
        # Spack's build environment when running tests
        vars_to_unset = ['CC', 'CXX', 'F77', 'FC']

        with spack.util.environment.preserve_environment(*vars_to_unset):
            # Delete temporarily the environment variables that point
            # to Spack compiler wrappers
            for v in vars_to_unset:
                del os.environ[v]
            # Check that mkoctfile outputs the expected value for CC
            cc = mkoctfile('-p', 'CC', output=str)
            msg = "mkoctfile didn't output the expected CC compiler"
            assert self.compiler.cc in cc, msg

            # Try to compile an Octave extension
            shutil.copy(helloworld_cc, tmp_dir)
            with working_dir(tmp_dir):
                mkoctfile('helloworld.cc')

    def configure_args(self):
        # See
        # https://github.com/macports/macports-ports/blob/master/math/octave/
        # https://github.com/Homebrew/homebrew-science/blob/master/octave.rb

        spec = self.spec
        config_args = []

        # Required dependencies
        config_args.extend([
            "--with-blas=%s" % spec['blas'].libs.ld_flags,
            "--with-lapack=%s" % spec['lapack'].libs.ld_flags
        ])

        # Strongly recommended dependencies
        if '+readline' in spec:
            config_args.append('--enable-readline')
        else:
            config_args.append('--disable-readline')

        # Optional dependencies
        if '+arpack' in spec:
            sa = spec['arpack-ng']
            config_args.extend([
                "--with-arpack-includedir=%s" % sa.prefix.include,
                "--with-arpack-libdir=%s"     % sa.prefix.lib
            ])
        else:
            config_args.append("--without-arpack")

        if '+curl' in spec:
            config_args.extend([
                "--with-curl-includedir=%s" % spec['curl'].prefix.include,
                "--with-curl-libdir=%s"     % spec['curl'].prefix.lib
            ])
        else:
            config_args.append("--without-curl")

        if '+fftw' in spec:
            config_args.extend([
                "--with-fftw3-includedir=%s"  % spec['fftw'].prefix.include,
                "--with-fftw3-libdir=%s"      % spec['fftw'].prefix.lib,
                "--with-fftw3f-includedir=%s" % spec['fftw'].prefix.include,
                "--with-fftw3f-libdir=%s"     % spec['fftw'].prefix.lib
            ])
        else:
            config_args.extend([
                "--without-fftw3",
                "--without-fftw3f"
            ])

        if '+fltk' in spec:
            config_args.extend([
                "--with-fltk-prefix=%s"      % spec['fltk'].prefix,
                "--with-fltk-exec-prefix=%s" % spec['fltk'].prefix
            ])
        else:
            config_args.append("--without-fltk")

        if '+glpk' in spec:
            config_args.extend([
                "--with-glpk-includedir=%s" % spec['glpk'].prefix.include,
                "--with-glpk-libdir=%s"     % spec['glpk'].prefix.lib
            ])
        else:
            config_args.append("--without-glpk")

        if '+magick' in spec:
            config_args.append("--with-magick=%s"
                                % spec['imagemagick'].prefix.lib)
        else:
            config_args.append("--without-magick")

        if '+hdf5' in spec:
            config_args.extend([
                "--with-hdf5-includedir=%s" % spec['hdf5'].prefix.include,
                "--with-hdf5-libdir=%s"     % spec['hdf5'].prefix.lib
            ])
        else:
            config_args.append("--without-hdf5")

        if '+jdk' in spec:
            config_args.extend([
                "--with-java-homedir=%s"    % spec['java'].home,
                "--with-java-includedir=%s" % spec['java'].home.include,
                "--with-java-libdir=%s"     % spec['java'].libs.directories[0]
            ])
        else:
            config_args.append("--disable-java")

        if '~opengl' and '~fltk' in spec:
            config_args.extend([
                "--without-opengl",
                "--without-framework-opengl"
            ])
        # TODO:  opengl dependency and package is missing?

        if '+qhull' in spec:
            config_args.extend([
                "--with-qhull-includedir=%s" % spec['qhull'].prefix.include,
                "--with-qhull-libdir=%s"     % spec['qhull'].prefix.lib
            ])
        else:
            config_args.append("--without-qhull")

        if '+qrupdate' in spec:
            config_args.extend([
                "--with-qrupdate-includedir=%s"
                % spec['qrupdate'].prefix.include,
                "--with-qrupdate-libdir=%s"     % spec['qrupdate'].prefix.lib
            ])
        else:
            config_args.append("--without-qrupdate")

        if '+zlib' in spec:
            config_args.extend([
                "--with-z-includedir=%s" % spec['zlib'].prefix.include,
                "--with-z-libdir=%s"     % spec['zlib'].prefix.lib
            ])
        else:
            config_args.append("--without-z")

        # If 64-bit BLAS is used:
        if (spec.satisfies('^openblas+ilp64') or
            spec.satisfies('^intel-mkl+ilp64') or
            spec.satisfies('^intel-parallel-studio+mkl+ilp64')):
            config_args.append('F77_INTEGER_8_FLAG=-fdefault-integer-8')

        # Use gfortran calling-convention %fj
        if spec.satisfies('%fj'):
            config_args.append('--enable-fortran-calling-convention=gfortran')

        return config_args

    # ========================================================================
    # Set up environment to make install easy for Octave extensions.
    # ========================================================================

    def setup_dependent_package(self, module, dependent_spec):
        """Called before Octave modules' install() methods.

        In most cases, extensions will only need to have one line:
            octave('--eval', 'pkg install %s' % self.stage.archive_file)
        """
        # Octave extension builds can have a global Octave executable function
        module.octave = Executable(join_path(self.spec.prefix.bin, 'octave'))

	`
	result, err := Decode(spackPackage)
	if err != nil {
		t.Fatal(err)
	}

	output, _ := PatchVersion(result, spackPackage)
	if output != spackPackage {
		t.Error(output)
		t.Error(spackPackage)
	}
}
