//go:build gofuzz
// +build gofuzz

package fuzz

// Fuzz is a fuzz test.
func Fuzz(data []byte) int { _ = "STUB: not implemented"; return 0 }

// ignore const-input error for OSS-Fuzz
