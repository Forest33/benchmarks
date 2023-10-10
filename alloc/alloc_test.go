package alloc

import (
	"math/rand"
	"testing"
)

const sliceSize = 1500

var testSlice []byte

func init() {
	testSlice = make([]byte, sliceSize)
	for i := 0; i < len(testSlice); i++ {
		testSlice[i] = byte(rand.Intn(256))
	}
}

// go:noinline
func allocFilter(in []byte) []byte {
	out := make([]byte, 0, sliceSize)
	for _, v := range in {
		if v <= 100 {
			out = append(out, v)
		}
	}
	return out
}

// go:noinline
func noAllocFilter(in []byte) []byte {
	out := in[:0]
	for _, v := range in {
		if v <= 100 {
			out = append(out, v)
		}
	}
	return out
}

func BenchmarkAllocFilter(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := allocFilter(testSlice)
		_ = out
	}
}

func BenchmarkNoAllocFilter(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := noAllocFilter(testSlice)
		_ = out
	}
}
