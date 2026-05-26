package main

/*
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef struct {
	const char *mediatype;
	const char *data;
	int32_t cssPrecision;
	int32_t cssVersion;
	bool htmlKeepComments;
	bool htmlKeepConditionalComments;
	bool htmlKeepDefaultAttrvals;
	bool htmlKeepDocumentTags;
	bool htmlKeepEndTags;
	bool htmlKeepQuotes;
	bool htmlKeepSpecialComments;
	bool htmlKeepWhitespace;
	bool jsKeepVarNames;
	int32_t jsPrecision;
	int32_t jsVersion;
	bool jsonKeepNumbers;
	int32_t jsonPrecision;
	bool svgKeepComments;
	int32_t svgPrecision;
	bool xmlKeepWhitespace;
} MinifyOptions;

typedef struct {
	char *error;
	char *data;
} MinifyResult;
*/
import "C"
import (
	"regexp"
	"unsafe"

	"github.com/tdewolff/minify/v2"
)

type minifyOptions struct {
	Type                        string
	Data                        string
	CSSPrecision                int
	CSSVersion                  int
	HTMLKeepComments            bool
	HTMLKeepConditionalComments bool
	HTMLKeepDefaultAttrvals     bool
	HTMLKeepDocumentTags        bool
	HTMLKeepEndTags             bool
	HTMLKeepQuotes              bool
	HTMLKeepSpecialComments     bool
	HTMLKeepWhitespace          bool
	JSKeepVarNames              bool
	JSPrecision                 int
	JSVersion                   int
	JSONKeepNumbers             bool
	JSONPrecision               int
	SVGKeepComments             bool
	SVGPrecision                int
	XMLKeepWhitespace           bool
}

var (
	jsMediatypePattern   = regexp.MustCompile("^(application|text)/(x-)?(java|ecma|j|live)script(1\\.[0-5])?$|^module$")
	jsonMediatypePattern = regexp.MustCompile("[/+]json$")
	xmlMediatypePattern  = regexp.MustCompile("[/+]xml$")
)

func parseOptions(opts *C.MinifyOptions) (minifyOptions, error) {
	_ = "STUB: not implemented"
	return *new(minifyOptions), nil
}

func newMinifier(opts minifyOptions) (*minify.M, error) { _ = "STUB: not implemented"; return nil, nil }

func resolveType(t string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func setResult(out *C.MinifyResult, err error, data string) { _ = "STUB: not implemented"; return }

//export Minify
func Minify(cOptions *C.MinifyOptions, cResult *C.MinifyResult) { _ = "STUB: not implemented"; return }

//export FreeCString
func FreeCString(ptr unsafe.Pointer) { _ = "STUB: not implemented"; return }

func main() {}
