package main

import (
	"io/fs"
)

func NewFS() fs.FS { _ = "STUB: not implemented"; return *new(fs.FS) }

type dirFS string

func (dir dirFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

func (dir dirFS) Stat(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}
