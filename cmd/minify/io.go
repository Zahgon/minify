package main

import (
	"io"
	"os"
)

// IsDir returns whether the path is a directory.
func IsDir(dir string) bool { _ = "STUB: not implemented"; return false }

// SameFile returns true if the two file paths specify the same path.
// While Linux is case-preserving case-sensitive (and therefore a string comparison will work),
// Windows is case-preserving case-insensitive; we use os.SameFile() to work cross-platform.
func SameFile(filename1 string, filename2 string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func openInputFile(input string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func openInputFiles(filenames []string, sep []byte) (*concatFileReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openOutputFile(output string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func createSymlink(input, output string) error { _ = "STUB: not implemented"; return nil }

type concatFileReader struct {
	filenames []string
	sep       []byte
	opener    func(string) (io.ReadCloser, error)

	cur     io.ReadCloser
	sepLeft int
}

func newConcatFileReader(filenames []string, opener func(string) (io.ReadCloser, error), sep []byte) (*concatFileReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *concatFileReader) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	// write remaining separator
	return 0, nil
}

// current reader is finished, load in the new reader

// if previous read returned (0, io.EOF), read from the new reader

func (r *concatFileReader) writeSep(p []byte) int { _ = "STUB: not implemented"; return 0 }

func (r *concatFileReader) Close() error { _ = "STUB: not implemented"; return nil }
