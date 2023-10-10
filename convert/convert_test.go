package main

import (
	"math/rand"
	"testing"
	"unsafe"
)

const sliceSize = 1500

var (
	benchSliceConvert  []byte
	benchStringConvert string
)

func init() {
	benchSliceConvert = make([]byte, sliceSize)
	for i := 0; i < len(benchSliceConvert); i++ {
		benchSliceConvert[i] = byte(rand.Intn(256))
	}
	benchStringConvert = string(benchSliceConvert)
}

// go:noinline
func byteSliceToString(in []byte) string {
	if len(in) == 0 {
		return ""
	}
	return string(in)
}

// go:noinline
func stringToByteSlice(in string) []byte {
	if len(in) == 0 {
		return nil
	}
	return []byte(in)
}

// go:noinline
func byteSliceToStringUnsafe(in []byte) string {
	if len(in) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(in), len(in))
}

// go:noinline
func stringToByteSliceUnsafe(in string) []byte {
	if in == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(in), len(in))
}

func BenchmarkBytesToString(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := byteSliceToString(benchSliceConvert)
		_ = out
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := stringToByteSlice(benchStringConvert)
		_ = out
	}
}

func BenchmarkBytesToStringUnsafe(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := byteSliceToStringUnsafe(benchSliceConvert)
		_ = out
	}
}

func BenchmarkStringToBytesUnsafe(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := stringToByteSliceUnsafe(benchStringConvert)
		_ = out
	}
}
