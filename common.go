package minify

import (
	"github.com/tdewolff/parse/v2"
)

var (
	textMimeBytes     = []byte("text/plain")
	charsetASCIIBytes = []byte("charset=us-ascii")
	dataBytes         = []byte("data:")
	base64Bytes       = []byte(";base64")
)

// Epsilon is the closest number to zero that is not considered to be zero.
var Epsilon = 0.00001

// Mediatype minifies a given mediatype by removing all whitespace and lowercasing all parts except strings (which may be case sensitive).
func Mediatype(b []byte) []byte { _ = "STUB: not implemented"; return nil }

// ToLower may otherwise slow down minification greatly

// DataURI minifies a data URI and calls a minifier by the specified mediatype. Specifications: https://www.ietf.org/rfc/rfc2397.txt.
func DataURI(m *M, dataURI []byte) []byte { _ = "STUB: not implemented"; return nil }

// must start with semicolon and be followed by end of mediatype or semicolon

// MaxInt is the maximum value of int.
const MaxInt = int(^uint(0) >> 1)

// MinInt is the minimum value of int.
const MinInt = -MaxInt - 1

// Decimal minifies a given byte slice containing a decimal and removes superfluous characters. It differs from Number in that it does not parse exponents.
// It does not parse or output exponents. prec is the number of significant digits. When prec is zero it will keep all digits. Only digits after the dot can be removed to reach the number of significant digits. Very large number may thus have more significant digits.
func Decimal(num []byte, prec int) []byte { _ = "STUB: not implemented"; return nil }

// omit first + and register mantissa start and end, whether it's negative and the exponent

// trim leading zeros but leave at least one digit

// trim trailing zeros

// apply precision

// include dot
// for numbers like .012

// process either an increase from a lesser significant decimal (>= 5)
// or remove trailing zeros after the dot, or both

// no-op

// end inc for integer

// Number minifies a given byte slice containing a number and removes superfluous characters.
func Number(num []byte, prec int) []byte { _ = "STUB: not implemented"; return nil }

// omit first + and register mantissa start and end, whether it's negative and the exponent

// range checks for when int is 32 bit

// trim leading zeros but leave at least one digit

// trim trailing zeros

// apply precision
//&& (dot <= start+prec || start+prec+1 < dot || 0 < origExp) { // don't minify 9 to 10, but do 999 to 1e3 and 99e1 to 1e3

// for numbers like .012

// for numbers where precision will include the dot

// do not minify 9=>10 or 99=>100 or 9e1=>1e2 (but 90), but 999=>1e3 and 99e1=>1e3

// process either an increase from a lesser significant decimal (>= 5)
// and remove trailing zeros

// no-op

// single digit left

// n is the number of significant digits
// normExp would be the exponent if it were normalised (0.1 <= f < 1)

// no number before exponent

// exponent overflow

// intExp would be the exponent if it were an integer

// there are three cases to consider when printing the number
// case 1: without decimals and with a positive exponent (large numbers: 5e4)
// case 2: with decimals and with a negative exponent (small numbers with many digits: .123456e-4)
// case 3: with decimals and without an exponent (around zero: 5.6)
// case 4: without decimals and with a negative exponent (small numbers: 123456e-9)

// case 1: print number with positive exponent

// remove dot, either from the front or copy the smallest part

// case 2: print normalized number (0.1 <= f < 1)

// case 3: print number without exponent

// place dot at the front, adding zeroes after the dot

// copy original digits after the dot towards the end

// copy original digits before the dot towards the end

// copy original digits before the dot towards the end

// place dot in the middle of the number

// when input has no dot in it

// when there are zeroes after the dot

// move digits between dot and newDot towards the end

// case 4: print number with negative exponent
// find new end, considering moving numbers to the front, removing the dot and increasing the length of the exponent

// it saves space to convert the decimal to an integer and decrease the exponent

// it does not save space and will panic, so we revert to the original representation

func UpdateErrorPosition(err error, input *parse.Input, offset int) error {
	_ = "STUB: not implemented"
	return nil
}
