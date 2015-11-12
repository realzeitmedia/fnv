package fnv

// Inline variant of hash/fnv's fnv64a.

const (
	offset64 = 14695981039346656037
	prime64  = 1099511628211
)

// New initializies a new fnv64a hash value.
func New() uint64 {
	return offset64
}

// Add adds a string to a fnv64a hash value, returning the updated hash.
func Add(h uint64, s string) uint64 {
	for i := 0; i < len(s); i++ {
		h = AddByte(h, s[i])
	}
	return h
}

// AddByte adds a byte to a fnv64a hash value, returning the updated hash.
func AddByte(h uint64, b byte) uint64 {
	h ^= uint64(b)
	h *= prime64
	return h
}

// AddBytes adds a byteslice to a fnv64a hash value, returning the updated hash.
func AddBytes(h uint64, bs []byte) uint64 {
	for _, b := range bs {
		h = AddByte(h, b)
	}
	return h
}
