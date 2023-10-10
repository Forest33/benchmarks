# Serialization

```
JSON:		7863 bytes
Msgpack:	7468 bytes
Protobuf:	3657 bytes

goos: linux
goarch: amd64
pkg: github.com/forest33/benchmarks/serialize
cpu: 12th Gen Intel(R) Core(TM) i9-12900H

BenchmarkMarshalJson
BenchmarkMarshalJson-20          	   62406	     16879 ns/op	    8222 B/op	       2 allocs/op
BenchmarkUnmarshalJson
BenchmarkUnmarshalJson-20        	   10000	    102214 ns/op	   16200 B/op	     761 allocs/op
BenchmarkMarshalJsonIter
BenchmarkMarshalJsonIter-20      	   90187	     13339 ns/op	    8224 B/op	       2 allocs/op
BenchmarkUnmarshalJsonIter
BenchmarkUnmarshalJsonIter-20    	   43840	     27432 ns/op	   12913 B/op	     519 allocs/op
BenchmarkMarshalMsgpack
BenchmarkMarshalMsgpack-20       	   53521	     22488 ns/op	   16404 B/op	      10 allocs/op
BenchmarkUnmarshalMsgpack
BenchmarkUnmarshalMsgpack-20     	   19858	     60255 ns/op	   19710 B/op	     760 allocs/op
BenchmarkMarshalProtobuf
BenchmarkMarshalProtobuf-20      	  157666	      7854 ns/op	    4096 B/op	       1 allocs/op
BenchmarkUnmarshalProtobuf
BenchmarkUnmarshalProtobuf-20    	   71700	     16893 ns/op	   16752 B/op	     457 allocs/op
```