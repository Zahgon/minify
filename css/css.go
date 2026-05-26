// Package css minifies CSS3 following the specifications at http://www.w3.org/TR/css-syntax-3/.
package css

import (
	"io"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/parse/v2/css"
)

var (
	spaceBytes        = []byte(" ")
	colonBytes        = []byte(":")
	semicolonBytes    = []byte(";")
	commaBytes        = []byte(",")
	leftBracketBytes  = []byte("{")
	rightBracketBytes = []byte("}")
	rightParenBytes   = []byte(")")
	urlBytes          = []byte("url(")
	varBytes          = []byte("var(")
	zeroBytes         = []byte("0")
	oneBytes          = []byte("1")
	transparentBytes  = []byte("transparent")
	blackBytes        = []byte("#0000")
	initialBytes      = []byte("initial")
	noneBytes         = []byte("none")
	autoBytes         = []byte("auto")
	leftBytes         = []byte("left")
	topBytes          = []byte("top")
	n400Bytes         = []byte("400")
	n700Bytes         = []byte("700")
	n50pBytes         = []byte("50%")
	n100pBytes        = []byte("100%")
	repeatXBytes      = []byte("repeat-x")
	repeatYBytes      = []byte("repeat-y")
	importantBytes    = []byte("!important")
	dataSchemeBytes   = []byte("data:")
)

type cssMinifier struct {
	m *minify.M
	w io.Writer
	p *css.Parser
	o *Minifier

	tokenBuffer []Token
	tokensLevel int
}

////////////////////////////////////////////////////////////////

// Minifier is a CSS minifier.
type Minifier struct {
	Precision    int // number of significant digits
	newPrecision int // precision for new numbers
	Inline       bool
	Version      int
}

// Minify minifies CSS data, it reads from r and writes to w.
func Minify(m *minify.M, w io.Writer, r io.Reader, params map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Token is a parsed token with extra information for functions.
type Token struct {
	css.TokenType
	Data       []byte
	Args       []Token // only filled for functions
	Fun, Ident Hash    // only filled for functions and identifiers respectively
}

func (t Token) String() string { _ = "STUB: not implemented"; return "" }

// Equal returns true if both tokens are equal.
func (t Token) Equal(t2 Token) bool { _ = "STUB: not implemented"; return false }

// IsZero return true if a dimension, percentage, or number token is zero.
func (t Token) IsZero() bool {
	_ = "STUB: not implemented"
	// as each number is already minified, starting with a zero means it is zero
	return false
}

// IsLength returns true if the token is a length.
func (t Token) IsLength() bool { _ = "STUB: not implemented"; return false }

// IsLengthPercentage returns true if the token is a length or percentage token.
func (t Token) IsLengthPercentage() bool { _ = "STUB: not implemented"; return false }

////////////////////////////////////////////////////////////////

// Minify minifies CSS data, it reads from r and writes to w.
func (o *Minifier) Minify(m *minify.M, w io.Writer, r io.Reader, params map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// minimum number of digits a double can represent exactly

func (c *cssMinifier) minifyGrammar() { _ = "STUB: not implemented"; return }

// write out the offending declaration (but save the semicolon)

func (c *cssMinifier) minifySelectors(values []css.Token) { _ = "STUB: not implemented"; return }

func (c *cssMinifier) parseFunction(values []css.Token) ([]Token, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// TODO: use ToHashFold

// TODO: use ToHashFold

func (c *cssMinifier) parseDeclaration(values []css.Token) []Token {
	_ = "STUB: not implemented"
	// Check if this is a simple list of values separated by whitespace or commas, otherwise we'll not be processing
	return nil
}

// TODO: use ToHashFold

// TODO: use ToHashFold

// update buffer size for memory reuse

func (c *cssMinifier) minifyDeclaration(property []byte, components []css.Token) {
	_ = "STUB: not implemented"
	return
}

// Strip !important from the component list, this will be added later separately

// Do not process complex values (eg. containing blocks or is not alternated between whitespace/commas and flat values

func (c *cssMinifier) writeFunction(args []Token) { _ = "STUB: not implemented"; return }

func (c *cssMinifier) writeDeclaration(values []Token, important bool) {
	_ = "STUB: not implemented"
	return
}

func (c *cssMinifier) minifyTokens(prop Hash, fun Hash, values []Token) []Token {
	_ = "STUB: not implemented"
	return nil
}

// integers

// don't use exponents

// don't use exponents

// cut dimension for zero value, TODO: don't hardcode check for Flex and remove the dimension in minifyDimension

// can overflow

// can overflow

// only minify color if fully opaque

// 0%, 20%, 40%, 60%, 80% and 100% can be represented exactly as, 51, 102, 153, 204, and 255 respectively

func (c *cssMinifier) minifyProperty(prop Hash, values []Token) []Token {
	_ = "STUB: not implemented"
	// limit maximum to prevent slow recursions (e.g. for background's append)
	return nil
}

// must contain atleast font-size and font-family
// the font-families are separated by commas and are at the end of font
// get index for last token before font family names

// identifier before first comma is a font-family

// advance i while still at font-families when they contain spaces but no quotes
// i cannot be 0, font-family must be prepended by font-size

// inherit, initial and unset are followed by an IdentToken/StringToken, so must be for font-size

// font-family minified in place

// fix for IE9, IE10, IE11: font name starting with `-` is not recognized

// line-height

// font-size

// if len is zero, it contains two consecutive spaces

// loop over comma-separated lists

// minify background-size and lowercase all identifiers

// background-size consists of either [<length-percentage> | auto | cover | contain] or [<length-percentage> | auto]{2}
// we can only minify the latter

// remove background-size if it is '/ auto' after minifying the property

// remove background-size if it is '/ auto'

// minify all other values
// position of background-origin that is padding-box

// background-origin

// background-clip

// further minify background-position and background-size combination

// loop over comma-separated lists

// loop over comma-separated lists

// loop over comma-separated lists

// remove zero offsets

// position of second set of horizontal/vertical values

// change right or bottom with percentage offset to left or top respectively

// removing zero offsets in the previous loop might make it eligible for the next loop

// we can't make this smaller, and converting to a number will break it
// (https://github.com/tdewolff/minify/issues/221#issuecomment-415419918)

// if it's a vertical position keyword, swap it with the next element
// since otherwise converted number positions won't be valid anymore
// (https://github.com/tdewolff/minify/issues/221#issue-353067229)

// transform keywords to lengths|percentages

// loop over comma-separated lists

// TODO: minify better (can be comma separated list)

// remove <flex-basis> if it is zero

// remove <flex-shrink> and <flex-basis> if they are 1 and 0 respectively

// remove auto to write 2-value syntax of <flex-grow> <flex-shrink>

// sort and remove overlapping ranges

// next range is fully contained in the current range

// next range is partially covering the current range

func minifyColor(value Token) Token { _ = "STUB: not implemented"; return *new(Token) }

// from working draft Color Module Level 4

func minifyNumberPercentage(value Token) Token {
	_ = "STUB: not implemented"
	// assumes input already minified
	return *new(Token)
}

func minifyLengthPercentage(value Token) Token { _ = "STUB: not implemented"; return *new(Token) }

// remove dimension for zero value

func (c *cssMinifier) minifyDimension(value Token) (Token, []byte) {
	_ = "STUB: not implemented"
	// TODO: add check for zero value
	return *new(Token), nil
}

// don't use exponents

// TODO: optimize
//if value.TokenType == css.DimensionToken {
//	// TODO: reverse; parse dim not number
//	n := parse.Number(value.Data)
//	num := value.Data[:n]
//	dim = value.Data[n:]
//	parse.ToLower(dim)

//	if c.o.Version<=2 {
//		num = minify.Decimal(num, c.o.Precision) // don't use exponents
//	} else {
//		num = minify.Number(num, c.o.Precision)
//	}

//	// change dimension to compress number
//	h := ToHash(dim)
//	if h == Px || h == Pt || h == Pc || h == In || h == Mm || h == Cm || h == Q || h == Deg || h == Grad || h == Rad || h == Turn || h == S || h == Ms || h == Hz || h == Khz || h == Dpi || h == Dpcm || h == Dppx {
//		d, _ := strconv.ParseFloat(string(num), 64) // can never fail
//		var dimensions []Hash
//		var multipliers []float64
//		switch h {
//		case Px:
//			//dimensions = []Hash{In, Cm, Pc, Mm, Pt, Q}
//			//multipliers = []float64{0.010416666666666667, 0.026458333333333333, 0.0625, 0.26458333333333333, 0.75, 1.0583333333333333}
//			dimensions = []Hash{In, Pc, Pt}
//			multipliers = []float64{0.010416666666666667, 0.0625, 0.75}
//		case Pt:
//			//dimensions = []Hash{In, Cm, Pc, Mm, Px, Q}
//			//multipliers = []float64{0.013888888888888889, 0.035277777777777778, 0.083333333333333333, 0.35277777777777778, 1.3333333333333333, 1.4111111111111111}
//			dimensions = []Hash{In, Pc, Px}
//			multipliers = []float64{0.013888888888888889, 0.083333333333333333, 1.3333333333333333}
//		case Pc:
//			//dimensions = []Hash{In, Cm, Mm, Pt, Px, Q}
//			//multipliers = []float64{0.16666666666666667, 0.42333333333333333, 4.2333333333333333, 12.0, 16.0, 16.933333333333333}
//			dimensions = []Hash{In, Pt, Px}
//			multipliers = []float64{0.16666666666666667, 12.0, 16.0}
//		case In:
//			//dimensions = []Hash{Cm, Pc, Mm, Pt, Px, Q}
//			//multipliers = []float64{2.54, 6.0, 25.4, 72.0, 96.0, 101.6}
//			dimensions = []Hash{Pc, Pt, Px}
//			multipliers = []float64{6.0, 72.0, 96.0}
//		case Cm:
//			//dimensions = []Hash{In, Pc, Mm, Pt, Px, Q}
//			//multipliers = []float64{0.39370078740157480, 2.3622047244094488, 10.0, 28.346456692913386, 37.795275590551181, 40.0}
//			dimensions = []Hash{Mm, Q}
//			multipliers = []float64{10.0, 40.0}
//		case Mm:
//			//dimensions = []Hash{In, Cm, Pc, Pt, Px, Q}
//			//multipliers = []float64{0.039370078740157480, 0.1, 0.23622047244094488, 2.8346456692913386, 3.7795275590551181, 4.0}
//			dimensions = []Hash{Cm, Q}
//			multipliers = []float64{0.1, 4.0}
//		case Q:
//			//dimensions = []Hash{In, Cm, Pc, Pt, Px} // Q to mm is never smaller
//			//multipliers = []float64{0.0098425196850393701, 0.025, 0.059055118110236220, 0.70866141732283465, 0.94488188976377953}
//			dimensions = []Hash{Cm} // Q to mm is never smaller
//			multipliers = []float64{0.025}
//		case Deg:
//			//dimensions = []Hash{Turn, Rad, Grad}
//			//multipliers = []float64{0.0027777777777777778, 0.017453292519943296, 1.1111111111111111}
//			dimensions = []Hash{Turn, Grad}
//			multipliers = []float64{0.0027777777777777778, 1.1111111111111111}
//		case Grad:
//			//dimensions = []Hash{Turn, Rad, Deg}
//			//multipliers = []float64{0.0025, 0.015707963267948966, 0.9}
//			dimensions = []Hash{Turn, Deg}
//			multipliers = []float64{0.0025, 0.9}
//		case Turn:
//			//dimensions = []Hash{Rad, Deg, Grad}
//			//multipliers = []float64{6.2831853071795865, 360.0, 400.0}
//			dimensions = []Hash{Deg, Grad}
//			multipliers = []float64{360.0, 400.0}
//		case Rad:
//			//dimensions = []Hash{Turn, Deg, Grad}
//			//multipliers = []float64{0.15915494309189534, 57.295779513082321, 63.661977236758134}
//		case S:
//			dimensions = []Hash{Ms}
//			multipliers = []float64{1000.0}
//		case Ms:
//			dimensions = []Hash{S}
//			multipliers = []float64{0.001}
//		case Hz:
//			dimensions = []Hash{Khz}
//			multipliers = []float64{0.001}
//		case Khz:
//			dimensions = []Hash{Hz}
//			multipliers = []float64{1000.0}
//		case Dpi:
//			dimensions = []Hash{Dppx, Dpcm}
//			multipliers = []float64{0.010416666666666667, 0.39370078740157480}
//		case Dpcm:
//			//dimensions = []Hash{Dppx, Dpi}
//			//multipliers = []float64{0.026458333333333333, 2.54}
//			dimensions = []Hash{Dpi}
//			multipliers = []float64{2.54}
//		case Dppx:
//			//dimensions = []Hash{Dpcm, Dpi}
//			//multipliers = []float64{37.795275590551181, 96.0}
//			dimensions = []Hash{Dpi}
//			multipliers = []float64{96.0}
//		}
//		for i := range dimensions {
//			if dimensions[i] != h { //&& (d < 1.0) == (multipliers[i] > 1.0) {
//				b := strconvParse.AppendFloat([]byte{}, d*multipliers[i], -1)
//				if c.o.Version<=2 {
//					b = minify.Decimal(b, c.o.newPrecision) // don't use exponents
//				} else {
//					b = minify.Number(b, c.o.newPrecision)
//				}
//				newDim := []byte(dimensions[i].String())
//				if len(b)+len(newDim) < len(num)+len(dim) {
//					num = b
//					dim = newDim
//				}
//			}
//		}
//	}
//	value.Data = append(num, dim...)
//}
//return value, dim
