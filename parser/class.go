package parser

import (
	"errors"
	"fmt"
	"strings"
)

// ParseClass handles the specific parsing of a package class.
func (parser *Parser) ParseClass() (err error) {
	token := parser.scnr.Get()
	if !token.IsClass() {
		return errors.New("called ParseClass without the beginning token being a class defintion")
	}

	err = parser.scnr.Next()
	if err != nil {
		if err.Error() == "end of input string" {
			return nil
		}
		return err
	}

	fmt.Println("Class")
	stmt := strings.Split(strings.Trim(token.Data, "):"), "(")
	parser.result.Name = stmt[0]
	parser.result.PackageType = stmt[1]

	return nil
}
