// Package html minifies HTML5 following the specifications at http://www.w3.org/TR/html5/syntax.html.
package html

import (
	"io"

	"github.com/tdewolff/minify/v2"
)

var (
	gtBytes         = []byte(">")
	isBytes         = []byte("=")
	spaceBytes      = []byte(" ")
	doctypeBytes    = []byte("<!doctype html>")
	jsMimeBytes     = []byte("application/javascript")
	cssMimeBytes    = []byte("text/css")
	htmlMimeBytes   = []byte("text/html")
	svgMimeBytes    = []byte("image/svg+xml")
	formMimeBytes   = []byte("application/x-www-form-urlencoded")
	mathMimeBytes   = []byte("application/mathml+xml")
	xmlMimeBytes    = []byte("text/xml")
	dataSchemeBytes = []byte("data:")
	jsSchemeBytes   = []byte("javascript:")
	httpBytes       = []byte("http")
	radioBytes      = []byte("radio")
	onBytes         = []byte("on")
	textBytes       = []byte("text")
	noneBytes       = []byte("none")
	submitBytes     = []byte("submit")
	allBytes        = []byte("all")
	rectBytes       = []byte("rect")
	dataBytes       = []byte("data")
	getBytes        = []byte("get")
	autoBytes       = []byte("auto")
	oneBytes        = []byte("one")
	inlineParams    = map[string]string{"inline": "1"}
)

////////////////////////////////////////////////////////////////

var GoTemplateDelims = [2]string{"{{", "}}"}
var HandlebarsTemplateDelims = [2]string{"{{", "}}"}
var MustacheTemplateDelims = [2]string{"{{", "}}"}
var EJSTemplateDelims = [2]string{"<%", "%>"}
var ASPTemplateDelims = [2]string{"<%", "%>"}
var PHPTemplateDelims = [2]string{"<?", "?>"}

// Minifier is an HTML minifier.
type Minifier struct {
	KeepComments            bool
	KeepConditionalComments bool
	KeepSpecialComments     bool
	KeepDefaultAttrVals     bool
	KeepDocumentTags        bool
	KeepEndTags             bool
	KeepQuotes              bool
	KeepWhitespace          bool
	TemplateDelims          [2]string
}

// Minify minifies HTML data, it reads from r and writes to w.
func Minify(m *minify.M, w io.Writer, r io.Reader, params map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Minify minifies HTML data, it reads from r and writes to w.
func (o *Minifier) Minify(m *minify.M, w io.Writer, r io.Reader, _ map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// omit next warning

// if true the next leading space is omitted

// [if ...] is always 7 or more characters, [endif] is only encountered for downlevel-revealed
// see https://msdn.microsoft.com/en-us/library/ms537512(v=vs.85).aspx#syntax
// downlevel-hidden

// malformed

// downlevel-revealed or short downlevel-hidden

// SSI tags

// omitSpace = true after block element

// whitespace removal; trim left

// whitespace removal; trim right

// trim if EOF, text token with leading whitespace or block token

// prevent long execution time with many following tokens

// stop looking when text encountered

// remove when followed by a block tag

// ignore empty script and style tags

// do not minify content of <style amp-boilerplate>

// EndTagToken

// remove superfluous tags, except for html, head and body tags when KeepDocumentTags is set

// omit end tags

// continue if text token is empty or whitespace

// omit p end tag

// continue if text token

// omit optgroup end tag

// omit spaces after block elements

// skip text in select and optgroup tags

// omit spaces after block elements

// mitigate for-loop increase

// write attributes

// removed attribute

// don't minify attributes that contain templates

// omit empty attribute values

// default attribute values can be omitted

// CSS minifier for attribute inline code

// JS minifier for attribute inline code

// anchors are already handled

// use double quotes for RDFa attributes

// no quotes if possible, else prefer single or double depending on which occurs more often in value

// StartTagClose

// skip text in select and optgroup tags

// keep space after phrasing tags (<i>, <span>, ...) FontAwesome etc.
