# Compression

```
Without compression:    2606 bytes
LZ4                     2165 bytes
LZO:                    1926 bytes
ZSTD:                   1389 bytes
Gzip:                   1360 bytes
Snappy:                 1950 bytes

goos: linux
goarch: amd64
pkg: github.com/forest33/benchmarks/compression
cpu: 12th Gen Intel(R) Core(TM) i9-12900H

BenchmarkCompressLZ4
BenchmarkCompressLZ4-20         	  224455	      4932 ns/op	    2699 B/op	       1 allocs/op
BenchmarkDecompressLZ4
BenchmarkDecompressLZ4-20       	 1958022	       618.5 ns/op	    2688 B/op	       1 allocs/op

BenchmarkCompressLZO
BenchmarkCompressLZO-20         	  228266	      5213 ns/op	    5320 B/op	       8 allocs/op
BenchmarkDecompressLZO
BenchmarkDecompressLZO-20       	  356234	      3267 ns/op	    7600 B/op	       3 allocs/op

BenchmarkCompressZSTD
BenchmarkCompressZSTD-20        	   37692	     31049 ns/op	    3563 B/op	       1 allocs/op
BenchmarkDecompressZSTD
BenchmarkDecompressZSTD-20      	  225399	      4963 ns/op	    2689 B/op	       1 allocs/op

BenchmarkCompressGzip
BenchmarkCompressGzip-20        	    7101	    166212 ns/op	 1079964 B/op	      20 allocs/op
BenchmarkDecompressGzip
BenchmarkDecompressGzip-20      	   79329	     14660 ns/op	   46992 B/op	      22 allocs/op

BenchmarkCompressSnappy
BenchmarkCompressSnappy-20      	   55969	     19746 ns/op	  144797 B/op	      15 allocs/op
BenchmarkDecompressSnappy
BenchmarkDecompressSnappy-20    	  129518	      9123 ns/op	   85680 B/op	       9 allocs/op
```
