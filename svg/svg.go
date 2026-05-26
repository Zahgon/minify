// Package svg minifies SVG1.1 following the specifications at http://www.w3.org/TR/SVG11/.
package svg

import (
	"io"

	"github.com/tdewolff/minify/v2"
)

var (
	voidBytes        = []byte("/>")
	isBytes          = []byte("=")
	spaceBytes       = []byte(" ")
	cdataEndBytes    = []byte("]]>")
	zeroBytes        = []byte("0")
	svgStartTagBytes = []byte("<svg")
	cssMimeBytes     = []byte("text/css")
	noneBytes        = []byte("none")
	urlBytes         = []byte("url(")
	xmlnsBytes       = []byte("xmlns")
)

////////////////////////////////////////////////////////////////

// Minifier is an SVG minifier.
type Minifier struct {
	KeepComments   bool
	Precision      int // number of significant digits
	KeepNamespaces []string
	newPrecision   int // precision for new numbers
	inline         bool
}

// Minify minifies SVG data, it reads from r and writes to w.
func Minify(m *minify.M, w io.Writer, r io.Reader, params map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Minify minifies SVG data, it reads from r and writes to w.
func (o *Minifier) Minify(m *minify.M, w io.Writer, r io.Reader, params map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// minimum number of digits a double can represent exactly

// namespaces to keep

// skip attributes in namespace (eg. inkscape or sodipodi)

// skip empty tags

// data is nil when attribute has been removed

// TODO: inefficient, temporary measure

// skip attributes in namespace (eg. inkscape or sodipodi)

//parse.ToLower(val)

// } else if len(val) > 5 && bytes.Equal(val[:4], []byte("rgb(")) && val[len(val)-1] == ')' {
// TODO: handle rgb(x, y, z) and hsl(x, y, z)

// prefer single or double quotes depending on what occurs more often in value

// collapse empty tags to single void tag

func (o *Minifier) shortenDimension(b []byte) ([]byte, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// only percentage is length 1

////////////////////////////////////////////////////////////////

func printTag(w io.Writer, tb *TokenBuffer, tag Hash) { _ = "STUB: not implemented"; return }

func skipTag(tb *TokenBuffer) { _ = "STUB: not implemented"; return }
