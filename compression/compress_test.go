package compression

import (
	"fmt"
	"testing"

	"github.com/forest33/benchmarks/entity"
)

const (
	zstdCompressionLevel = 2
)

func init() {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	lz4, _ := cmp.CompressLZ4(entity.TextData)
	lzo, _ := cmp.CompressLZO(entity.TextData)
	zstd, _ := cmp.CompressZSTD(entity.TextData, zstdCompressionLevel)
	gzip, _ := cmp.CompressGzip(entity.TextData)
	snappy, _ := cmp.CompressSnappy(entity.TextData)

	fmt.Printf("Without compression:\t%d bytes\n", len(entity.TextData))
	fmt.Printf("LZ4:\t\t\t\t\t%d bytes\n", len(lz4))
	fmt.Printf("LZO:\t\t\t\t\t%d bytes\n", len(lzo))
	fmt.Printf("ZSTD:\t\t\t\t\t%d bytes\n", len(zstd))
	fmt.Printf("Gzip:\t\t\t\t\t%d bytes\n", len(gzip))
	fmt.Printf("Snappy:\t\t\t\t\t%d bytes\n", len(snappy))
}

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

func BenchmarkCompressGzip(b *testing.B) {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := cmp.CompressGzip(entity.TextData)
		_ = out
	}
}

func BenchmarkDecompressGzip(b *testing.B) {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	data, _ := cmp.CompressGzip(entity.TextData)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := cmp.DecompressGzip(data)
		_ = out
	}
}

func BenchmarkCompressSnappy(b *testing.B) {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := cmp.CompressSnappy(entity.TextData)
		_ = out
	}
}

func BenchmarkDecompressSnappy(b *testing.B) {
	cmp := New(&Config{PayloadSize: len(entity.TextData)})
	data, _ := cmp.CompressSnappy(entity.TextData)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := cmp.DecompressSnappy(data)
		_ = out
	}
}
