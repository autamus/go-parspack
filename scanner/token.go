package scanner

import "strings"

// Token is the structure containing a scanned token.
type Token struct {
	Data string
}

// IsComment returns if the token is a comment.
func (t *Token) IsComment() bool {
	return strings.HasPrefix(t.Data, "#")
}

// IsString returns if the current token begins a string.
func (t *Token) IsString() bool {
	return strings.HasPrefix(t.Data, `"""`) ||
		strings.HasPrefix(t.Data, `"`) ||
		strings.HasPrefix(t.Data, `'`) ||
		strings.HasSuffix(t.Data, `"""`) ||
		strings.HasSuffix(t.Data, `"`) ||
		strings.HasSuffix(t.Data, `'`)
}

// IsClass returns if the current token is begins a class definition.
func (t *Token) IsClass() bool {
	return strings.ToLower(t.Data) == "class"
}
