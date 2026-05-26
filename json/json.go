// Package json minifies JSON following the specifications at http://json.org/.
package json

import (
	"io"

	"github.com/tdewolff/minify/v2"
)

var (
	commaBytes     = []byte(",")
	colonBytes     = []byte(":")
	zeroBytes      = []byte("0")
	minusZeroBytes = []byte("-0")
)

////////////////////////////////////////////////////////////////

// Minifier is a JSON minifier.
type Minifier struct {
	Precision   int  // number of significant digits
	KeepNumbers bool // prevent numbers from being minified
}

// Minify minifies JSON data, it reads from r and writes to w.
func Minify(m *minify.M, w io.Writer, r io.Reader, params map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Minify minifies JSON data, it reads from r and writes to w.
func (o *Minifier) Minify(_ *minify.M, w io.Writer, r io.Reader, _ map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}
