package compression

import (
	"testing"

	"github.com/forest33/benchmarks/entity"
)

const (
	zstdCompressionLevel = 2
)

func BenchmarkCompressLZ4(b *testing.B) {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := cmp.CompressLZ4(entity.TextData)
		_ = out
	}
}

func BenchmarkDecompressLZ4(b *testing.B) {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	data, _ := cmp.CompressLZ4(entity.TextData)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := cmp.DecompressLZ4(data)
		_ = out
	}
}

func BenchmarkCompressLZO(b *testing.B) {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := cmp.CompressLZO(entity.TextData)
		_ = out
	}
}

func BenchmarkDecompressLZO(b *testing.B) {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	data, _ := cmp.CompressLZO(entity.TextData)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := cmp.DecompressLZO(data)
		_ = out
	}
}

func BenchmarkCompressZSTD(b *testing.B) {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := cmp.CompressZSTD(entity.TextData, zstdCompressionLevel)
		_ = out
	}
}

func BenchmarkDecompressZSTD(b *testing.B) {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	data, _ := cmp.CompressZSTD(entity.TextData, zstdCompressionLevel)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := cmp.DecompressZSTD(data)
		_ = out
	}
}
