package main

import "C"
import (
	"github.com/tdewolff/minify/v2"
)

var m *minify.M

func init() {
	minifyConfig(nil, nil, 0)
}

func goBytes(str *C.char, length, capacity C.longlong) []byte {
	_ = "STUB: not implemented"
	// somehow address space on 32-bit system is smaller than 1<<30
	return nil
}

func goStringArray(carr **C.char, length C.longlong) []string {
	_ = "STUB: not implemented"
	return nil
}

//export minifyConfig
func minifyConfig(ckeys **C.char, cvals **C.char, length C.longlong) *C.char {
	_ = "STUB: not implemented"
	return nil
}

// also handles <?php

//export minifyString
func minifyString(cmediatype, cinput *C.char, input_length C.longlong, coutput *C.char, output_length *C.longlong) *C.char {
	_ = "STUB: not implemented"
	return nil
}

// copy
// +1 for NULL byte used in parser

//export minifyFile
func minifyFile(cmediatype, cinput, coutput *C.char) *C.char { _ = "STUB: not implemented"; return nil }

// copy

func main() {}
