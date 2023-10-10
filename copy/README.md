# Copying slices

```
goos: linux
goarch: amd64
pkg: github.com/forest33/benchmarks/copy
cpu: 12th Gen Intel(R) Core(TM) i9-12900H

BenchmarkClone
BenchmarkClone-20       	 6263290	       192.9 ns/op	    1536 B/op	       1 allocs/op
BenchmarkCopy
BenchmarkCopy-20        	 5264325	       191.7 ns/op	    1536 B/op	       1 allocs/op
BenchmarkCopyPool
BenchmarkCopyPool-20    	70305795	        19.72 ns/op	       0 B/op	       0 allocs/op
```