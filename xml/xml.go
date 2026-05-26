// Package xml minifies XML1.0 following the specifications at http://www.w3.org/TR/xml/.
package xml

import (
	"io"

	"github.com/tdewolff/minify/v2"
)

var (
	isBytes    = []byte("=")
	spaceBytes = []byte(" ")
	voidBytes  = []byte("/>")
)

////////////////////////////////////////////////////////////////

// Minifier is an XML minifier.
type Minifier struct {
	KeepWhitespace bool
}

// Minify minifies XML data, it reads from r and writes to w.
func Minify(m *minify.M, w io.Writer, r io.Reader, params map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Minify minifies XML data, it reads from r and writes to w.
func (o *Minifier) Minify(m *minify.M, w io.Writer, r io.Reader, _ map[string]string) error {
	_ = "STUB: not implemented"
	// on true the next text token must not start with a space
	return nil
}

// convert CDATA to regular text if smaller

// whitespace removal; trim left

// whitespace removal; trim right

// trim if EOF, text token with whitespace begin or block token

// this only happens when a comment, doctype, cdata startpi tag was in between
// remove if the text token starts with a whitespace

// prefer single or double quotes depending on what occurs more often in value

// collapse empty tags to single void tag
