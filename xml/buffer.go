package xml

import (
	"github.com/tdewolff/parse/v2/xml"
)

// Token is a single token unit with an attribute value (if given) and hash of the data.
type Token struct {
	xml.TokenType
	Data    []byte
	Text    []byte
	AttrVal []byte
}

// TokenBuffer is a buffer that allows for token look-ahead.
type TokenBuffer struct {
	l *xml.Lexer

	buf []Token
	pos int
}

// NewTokenBuffer returns a new TokenBuffer.
func NewTokenBuffer(l *xml.Lexer) *TokenBuffer { _ = "STUB: not implemented"; return nil }

func (z *TokenBuffer) read(t *Token) { _ = "STUB: not implemented"; return }

// Peek returns the ith element and possibly does an allocation.
// Peeking past an error will panic.
func (z *TokenBuffer) Peek(pos int) *Token { _ = "STUB: not implemented"; return nil }

// required peek length

// Shift returns the first element and advances position.
func (z *TokenBuffer) Shift() *Token { _ = "STUB: not implemented"; return nil }
