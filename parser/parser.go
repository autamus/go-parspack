package parser

import (
	"fmt"
	"strings"

	"github.com/autamus/go-parspack/pkg"
	"github.com/autamus/go-parspack/scanner"
)

// Parser is the collection responsible for decoding an input string into
// a package struct.
type Parser struct {
	scnr   scanner.Scanner
	result pkg.Package
}

// Init initializes a Parser with a scanner.
func (parser *Parser) Init(scnr scanner.Scanner) {
	parser.scnr = scnr
	parser.result = pkg.Package{}
}

// Parse decodes the information from a scanner into a package struct.
func (parser *Parser) Parse() (result pkg.Package, err error) {
	if !parser.scnr.hasNextLine() && !parser.scnr.hasNextOnLine() {
		return parser.result, nil
	}

	token := parser.scnr.get()

	switch {
	case strings.HasPrefix(token, "#") && parser.scnr.hasNextLine():
		parser.scnr.nextLine()
		parser.Parse()

	case strings.HasPrefix(token, `"""`):
		fmt.Println("Description")

	case token == "class":
		fmt.Println("Class")
		token = parser.scnr.get()
		if err != nil {
			return result, err
		}
		stmt := strings.Split(strings.Trim(token, "):"), "(")
		parser.result.Name = stmt[0]
		parser.result.PackageType = stmt[1]
	}

	err = parser.scnr.next()
	if err != nil {
		if err.Error() == "end of input string" {
			return parser.result, nil
		}
		return parser.result, err
	}
	return parser.Parse()
}
