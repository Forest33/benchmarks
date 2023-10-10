# Memory allocation

```
goos: linux
goarch: amd64
pkg: github.com/forest33/benchmarks/alloc
cpu: 12th Gen Intel(R) Core(TM) i9-12900H

BenchmarkAllocFilter
BenchmarkAllocFilter-20      	 1206176	       988.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoAllocFilter
BenchmarkNoAllocFilter-20    	 2026276	       590.4 ns/op	       0 B/op	       0 allocs/op
```