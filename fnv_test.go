package fnv_test

import (
	theirlib "hash/fnv"
	"testing"

	ourlib "github.com/realzeitmedia/fnv"
)

func TestString(t *testing.T) {
	for _, src := range []string{
		"",
		"hello world",
	} {
		h := theirlib.New64a()
		h.Write([]byte(src))
		theirs := h.Sum64()

		ours := ourlib.New()
		ours = ourlib.Add(ours, src)

		if have, want := ours, theirs; have != want {
			t.Errorf("have %d, want %d. Case %q", src)
		}
	}
}
