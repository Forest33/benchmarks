# Serialization

```
JSON:		7904 bytes
Msgpack:	7523 bytes
Protobuf:	3639 bytes
Gob:		3739 bytes

goos: linux
goarch: amd64
pkg: github.com/forest33/benchmarks/serialize
cpu: 12th Gen Intel(R) Core(TM) i9-12900H

BenchmarkMarshalJson
BenchmarkMarshalJson-20          	   69676	     17380 ns/op	    8222 B/op	       2 allocs/op
BenchmarkUnmarshalJson
BenchmarkUnmarshalJson-20        	   10000	    108958 ns/op	   16392 B/op	     778 allocs/op

BenchmarkMarshalJsonIter
BenchmarkMarshalJsonIter-20      	   88530	     14040 ns/op	    8223 B/op	       2 allocs/op
BenchmarkUnmarshalJsonIter
BenchmarkUnmarshalJsonIter-20    	   41410	     29590 ns/op	   13138 B/op	     529 allocs/op

BenchmarkMarshalMsgpack
BenchmarkMarshalMsgpack-20       	   50258	     24454 ns/op	   16404 B/op	      10 allocs/op
BenchmarkUnmarshalMsgpack
BenchmarkUnmarshalMsgpack-20     	   19470	     63243 ns/op	   19727 B/op	     756 allocs/op

BenchmarkMarshalGob
BenchmarkMarshalGob-20           	   96715	     12173 ns/op	    2136 B/op	      89 allocs/op
BenchmarkUnmarshalGob
BenchmarkUnmarshalGob-20         	 6014751	       194.5 ns/op	     920 B/op	       2 allocs/op

BenchmarkMarshalProtobuf
BenchmarkMarshalProtobuf-20      	  143959	      7633 ns/op	    4096 B/op	       1 allocs/op
BenchmarkUnmarshalProtobuf
BenchmarkUnmarshalProtobuf-20    	   66168	     17680 ns/op	   16936 B/op	     463 allocs/op
```