# Type conversion

```
goos: linux
goarch: amd64
pkg: github.com/forest33/benchmarks/convert
cpu: 12th Gen Intel(R) Core(TM) i9-12900H

BenchmarkBytesToString
BenchmarkBytesToString-20          	 6691567	       189.0 ns/op	    1536 B/op	       1 allocs/op
BenchmarkStringToBytes
BenchmarkStringToBytes-20          	 6358797	       196.5 ns/op	    1536 B/op	       1 allocs/op

BenchmarkBytesToStringUnsafe
BenchmarkBytesToStringUnsafe-20    	1000000000	         0.3335 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringToBytesUnsafe
BenchmarkStringToBytesUnsafe-20    	1000000000	         0.3344 ns/op	       0 B/op	       0 allocs/op
```