package scanner

import (
	"errors"
	"strings"
)

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

// nextLine sets the cursor to the beginning of the next line.
func (scnr *Scanner) nextLine() {
	scnr.cursor = 0
	scnr.lineIndex++

	scnr.currentLine = strings.Fields(scnr.lines[scnr.lineIndex])

	if len(scnr.currentLine) == 0 {
		scnr.nextLine()
	}
}

func (scnr *Scanner) get() (result string) {
	return scnr.currentLine[scnr.cursor]
}

func (scnr *Scanner) next() (err error) {
	switch {
	case scnr.hasNextOnLine():
		scnr.cursor++

	case scnr.hasNextLine():
		scnr.nextLine()

	default:
		return errors.New("end of input string")
	}
	return nil
}

func (scnr *Scanner) hasNextOnLine() bool {
	return scnr.cursor < len(scnr.currentLine)
}

func (scnr *Scanner) hasNextLine() bool {
	return scnr.lineIndex < len(scnr.lines)
}
