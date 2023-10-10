package main

import (
	"bytes"
	"math/rand"
	"slices"
	"testing"
)

const sliceSize = 1500

var benchSliceCopy []byte

func init() {
	benchSliceCopy = make([]byte, sliceSize)
	for i := 0; i < len(benchSliceCopy); i++ {
		benchSliceCopy[i] = byte(rand.Intn(256))
	}
	bPool = NewBytesPool()
}

var bPool *BytesPool

// go:noinline
func cloneData(in []byte) []byte {
	return slices.Clone(in)
}

// go:noinline
func copyData(in []byte) []byte {
	out := make([]byte, len(in))
	copy(out, in)
	return out
}

// go:noinline
func copyDataPool(in []byte) *bytes.Buffer {
	out := bPool.Get()
	out.Write(in)
	return out
}

func BenchmarkClone(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := cloneData(benchSliceCopy)
		_ = out
	}
}

func BenchmarkCopy(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := copyData(benchSliceCopy)
		_ = out
	}
}

func BenchmarkCopyPool(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := copyDataPool(benchSliceCopy)
		_ = out
		bPool.Release(out)
	}
}
