# Encryption

```
goos: linux
goarch: amd64
pkg: github.com/forest33/benchmarks/encryption
cpu: 12th Gen Intel(R) Core(TM) i9-12900H

BenchmarkAESECB_Encrypt
BenchmarkAESECB_Encrypt-20    	  601491	      1789 ns/op	    6146 B/op	       3 allocs/op
BenchmarkAESECB_Decrypt
BenchmarkAESECB_Decrypt-20    	  946488	      1244 ns/op	    2712 B/op	       2 allocs/op
BenchmarkAESGCM_Encrypt
BenchmarkAESGCM_Encrypt-20    	 1000000	      1007 ns/op	    2704 B/op	       2 allocs/op
BenchmarkAESGCM_Decrypt
BenchmarkAESGCM_Decrypt-20    	 1555808	       746.1 ns/op	    2712 B/op	       2 allocs/op
```