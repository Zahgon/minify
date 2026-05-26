//go:build !linux && !darwin && !netbsd && !solaris && !openbsd && !js && !wasm

package main

import "os"

var supportsGetOwnership = false

func getOwnership(info os.FileInfo) (int, int, bool) { _ = "STUB: not implemented"; return 0, 0, false }
