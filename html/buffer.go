package html

import (
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/html"
)

// Token is a single token unit with an attribute value (if given) and hash of the data.
type Token struct {
	html.TokenType
	Hash        Hash
	Data        []byte
	Text        []byte
	AttrVal     []byte
	Traits      traits
	Offset      int
	HasTemplate bool
}

// TokenBuffer is a buffer that allows for token look-ahead.
type TokenBuffer struct {
	r *parse.Input
	l *html.Lexer

	buf []Token
	pos int

	attrBuffer []*Token
}

// NewTokenBuffer returns a new TokenBuffer.
func NewTokenBuffer(r *parse.Input, l *html.Lexer) *TokenBuffer {
	_ = "STUB: not implemented"
	return nil
}

func (z *TokenBuffer) read(t *Token) { _ = "STUB: not implemented"; return }

// quotes will be readded in attribute loop if necessary

// zero if not exist

// Peek returns the ith element and possibly does an allocation.
// Peeking past an error will panic.
func (z *TokenBuffer) Peek(pos int) *Token { _ = "STUB: not implemented"; return nil }

// required peek length

// Shift returns the first element and advances position.
func (z *TokenBuffer) Shift() *Token { _ = "STUB: not implemented"; return nil }

// Attributes extracts the gives attribute hashes from a tag.
// It returns in the same order pointers to the requested token data or nil.
func (z *TokenBuffer) Attributes(hashes ...Hash) []*Token { _ = "STUB: not implemented"; return nil }
