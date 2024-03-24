package jlexer

import "fmt"

// LexerError implements the error interface and represents all possible errors that can be
// generated during parsing the JSON data.
type LexerError struct {
	Reason string
	Offset int
	Data   string

	// Fields that are filled in when verbose mode is enabled.
	Key              string
	IsUnsupportedKey bool
	IsInvalidValue   bool
}

func (l *LexerError) Error() string {
	return fmt.Sprintf("parse error: %s near offset %d of '%s'", l.Reason, l.Offset, l.Data)
}

func (l *LexerError) VerboseError() string {
	if l.IsUnsupportedKey {
		return fmt.Sprintf("'%s' is unsupported", l.Key)
	}
	if l.IsInvalidValue {
		return fmt.Sprintf("'%s' has incorrect value", l.Key)
	}
	return l.Error()
}
