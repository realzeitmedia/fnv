Simple Go FNV64a hash

This avoids allocations compared to hash/fnv from the standard library.

```
$ go test -bench=. -benchmem
PASS
BenchmarkStringsStdlib-4	20000000	        87.5 ns/op	      16 B/op	       1 allocs/op
BenchmarkStrings-4      	50000000	        33.0 ns/op	       0 B/op	       0 allocs/op
ok  	github.com/realzeitmedia/fnv	3.540s
```

See example_test.go for usage.

Installation:

    `go get github.com/realzeitmedia/fnv`

