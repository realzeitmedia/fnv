package fnv_test

import (
	theirlib "hash/fnv"
	"testing"

	ourlib "github.com/realzeitmedia/fnv"
)

const (
	src = "hello world"
)

func BenchmarkStringsStdlib(b *testing.B) {
	h := theirlib.New64a()
	for i := 0; i < b.N; i++ {
		h.Write([]byte(src))
	}
	_ = h.Sum64()
}

func BenchmarkStrings(b *testing.B) {
	h := ourlib.New()
	for i := 0; i < b.N; i++ {
		h = ourlib.Add(h, src)
	}
}
