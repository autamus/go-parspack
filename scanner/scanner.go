package scanner

import (
	"errors"
	"strings"
)

// Scanner is a structure containing the information of where the
// scanner cursor is located and the input text for the scanner.
type Scanner struct {
	lines       []string
	currentLine []string
	lineIndex   int
	cursor      int
}

// Init initializes the scanner with the given input string.
func (scnr *Scanner) Init(input string) {
	scnr.lineIndex = 0
	scnr.cursor = 0

	scnr.lines = strings.Split(input, "\n")
	scnr.currentLine = strings.Fields(scnr.lines[scnr.lineIndex])
}

// NextLine sets the cursor to the beginning of the next line.
func (scnr *Scanner) NextLine() {
	scnr.cursor = 0
	scnr.lineIndex++

	scnr.currentLine = strings.Fields(scnr.lines[scnr.lineIndex])

	if len(scnr.currentLine) == 0 {
		scnr.NextLine()
	}
}

// Get returns the current token from the scanner.
func (scnr *Scanner) Get() (result Token) {
	return Token{scnr.currentLine[scnr.cursor]}
}

// Next moves to the next token in the input text.
func (scnr *Scanner) Next() (err error) {
	switch {
	case scnr.HasNextOnLine():
		scnr.cursor++

	case scnr.HasNextLine():
		scnr.NextLine()

	default:
		return errors.New("end of input string")
	}
	return nil
}

// HasNextOnLine returns if the scanner contains another token on the same line.
func (scnr *Scanner) HasNextOnLine() bool {
	return scnr.cursor < len(scnr.currentLine)
}

// HasNextLine returns if the scanner contains another line of data.
func (scnr *Scanner) HasNextLine() bool {
	return scnr.lineIndex < len(scnr.lines)
}
