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
			t.Errorf("have %d, want %d. case %q", have, want, src)
		}
	}
}

func TestStrings(t *testing.T) {
	for _, src := range [][]string{
		{},
		{""},
		{"hello", "world"},
	} {
		h := theirlib.New64a()
		for _, s := range src {
			h.Write([]byte(s))
		}
		theirs := h.Sum64()

		ours := ourlib.New()
		for _, s := range src {
			ours = ourlib.Add(ours, s)
		}

		if have, want := ours, theirs; have != want {
			t.Errorf("have %d, want %d. case %q", have, want, src)
		}
	}
}

func TestByte(t *testing.T) {
	for _, src := range []string{
		"",
		"hello world",
	} {
		h := theirlib.New64a()
		h.Write([]byte(src))
		theirs := h.Sum64()

		ours := ourlib.New()
		for i := 0; i < len(src); i++ {
			ours = ourlib.AddByte(ours, src[i])
		}

		if have, want := ours, theirs; have != want {
			t.Errorf("have %d, want %d. case %q", have, want, src)
		}
	}
}
