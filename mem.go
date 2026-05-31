package gohelp

import "unsafe"

// Convert unsafe convert. Use carefully
func Convert[S, T any](source S) T {
	// #nosec G103
	return *(*T)(unsafe.Pointer(&source))
}

// Ptr get pointer of value
func Ptr[T any](v T) *T {
	return &v
}
