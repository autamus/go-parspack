package scanner

import (
	"regexp"
	"strings"
)

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
	strTest := regexp.MustCompile(`(^""".*$)|(^".*$)|(^'.*$)|(^[^"""]*"""[,|\)]*$)|(^[^"]*"[,|\)]*$)|(^[^']*'[,|\)]*$)`)
	return strTest.MatchString(t.Data)
}

// IsClass returns if the current token is begins a class definition.
func (t *Token) IsClass() bool {
	return strings.ToLower(t.Data) == "class"
}

// IsHomepage returns if the current token begins with a homepage variable declaration.
func (t *Token) IsHomepage() bool {
	return strings.ToLower(t.Data) == "homepage"
}

// IsURL returns if the current token begins with a valid URL header.
func (t *Token) IsURL() bool {
	return strings.HasPrefix(t.Data, "url")
}

//IsChecksum returns uf the current token begins with a valid checksum header.
func (t *Token) IsChecksum() bool {
	return strings.HasPrefix(strings.ToLower(t.Data), "sha")
}

// IsVersion returns if the current token begins with a valid version header.
func (t *Token) IsVersion() bool {
	return strings.HasPrefix(t.Data, "version(")
}

// IsDependency returns if the current token begins with a valid version header.
func (t *Token) IsDependency() bool {
	return strings.HasPrefix(strings.ToLower(t.Data), "depends_on(")
}

// IsFunction returns if the current token begins a function.
func (t *Token) IsFunction() bool {
	return strings.ToLower(t.Data) == "def"
}

// IsBranch returns if the current token defines a branch keyword.
func (t *Token) IsBranch() bool {
	return strings.HasPrefix(strings.ToLower(t.Data), "branch")
}

// IsBoolean returns if the current token is a boolean.
func (t *Token) IsBoolean() bool {
	return strings.HasPrefix(strings.ToLower(t.Data), "true") ||
		strings.HasPrefix(strings.ToLower(t.Data), "false")
}

// IsSubmodule returns if the current token is a submodule keyword.
func (t *Token) IsSubmodule() bool {
	return strings.HasPrefix(strings.ToLower(t.Data), "submodules")
}

// IsExpand returns if the current token is an expand keyword.
func (t *Token) IsExpand() bool {
	return strings.HasPrefix(strings.ToLower(t.Data), "expand")
}
