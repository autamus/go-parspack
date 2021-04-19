package pkg_test

import (
	"fmt"
	"testing"

	"github.com/DataDrake/cuppa/version"
	"github.com/autamus/go-parspack/pkg"
)

func TestAdd(t *testing.T) {
	original := pkg.Package{Name: "Go", Versions: []pkg.Version{{Value: version.NewVersion("1.16")}}}
	original.AddVersion(pkg.Version{Value: version.NewVersion("1.16.2")})
	fmt.Println(original)
	t.Fatal()
}
