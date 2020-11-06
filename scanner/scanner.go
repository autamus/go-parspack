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

// Init instantiates a new scanner with the input text.
func (s *Scanner) Init(input string) {
	s.lines = strings.Split(input, "\n")
	s.lineIndex = -1
	s.NextLine()
}

// Next moves the cursor to the next token.
func (s *Scanner) Next() (err error) {
	switch {
	case s.cursor < len(s.currentLine)-1:
		s.cursor++

	case s.lineIndex < len(s.lines)-1:
		err = s.NextLine()

	default:
		err = errors.New("end of scanner source")
	}
	return err
}

// NextLine moves the cursor to the next line if possible.
func (s *Scanner) NextLine() (err error) {
	var i int
	for i = s.lineIndex + 1; i < len(s.lines); i++ {
		line := strings.Fields(s.lines[i])
		if len(line) > 0 {
			s.lineIndex = i
			s.currentLine = line
			s.cursor = 0
			break
		}
	}
	if i >= len(s.lines) {
		return errors.New("end of scanner source")
	}
	return nil
}

// Peak grabs the current token.
func (s *Scanner) Peak() (result Token) {
	return Token{Data: s.currentLine[s.cursor]}
}
