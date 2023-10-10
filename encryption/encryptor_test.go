package encryptor

import (
	"strings"
	"testing"

	"github.com/forest33/benchmarks/entity"
)

var (
	key = strings.Repeat("0", 32)
)

func BenchmarkAESECB_Encrypt(b *testing.B) {
	enc := NewAESECB(key)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := enc.Encrypt(entity.TextData)
		l := enc.GetLength(len(entity.TextData))
		_ = out
		_ = l
	}
}

func BenchmarkAESECB_Decrypt(b *testing.B) {
	enc := NewAESECB(key)
	data, _ := enc.Encrypt(entity.TextData)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := enc.Decrypt(data)
		_ = out
	}
}

func BenchmarkAESGCM_Encrypt(b *testing.B) {
	enc := NewAESGCM(key)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := enc.Encrypt(entity.TextData)
		l := enc.GetLength(len(entity.TextData))
		_ = out
		_ = l
	}
}

func BenchmarkAESGCM_Decrypt(b *testing.B) {
	enc := NewAESGCM(key)
	data, _ := enc.Encrypt(entity.TextData)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := enc.Decrypt(data)
		_ = out
	}
}
