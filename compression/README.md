# Compression

```
goos: linux
goarch: amd64
pkg: github.com/forest33/benchmarks/compression
cpu: 12th Gen Intel(R) Core(TM) i9-12900H

BenchmarkCompressLZ4
BenchmarkCompressLZ4-20       	  237631	      4970 ns/op	    2695 B/op	       1 allocs/op
BenchmarkDecompressLZ4
BenchmarkDecompressLZ4-20     	 1892149	       627.3 ns/op	    2688 B/op	       1 allocs/op

BenchmarkCompressLZO
BenchmarkCompressLZO-20       	  229399	      5138 ns/op	    5320 B/op	       8 allocs/op
BenchmarkDecompressLZO
BenchmarkDecompressLZO-20     	  311006	      3344 ns/op	    7600 B/op	       3 allocs/op

BenchmarkCompressZSTD
BenchmarkCompressZSTD-20      	   38032	     30120 ns/op	    3555 B/op	       1 allocs/op
BenchmarkDecompressZSTD
BenchmarkDecompressZSTD-20    	  212612	      4980 ns/op	    2689 B/op	       1 allocs/op
```