package fnv_test

import (
	"fmt"

	"github.com/realzeitmedia/fnv"
)

func Example() {
	sum := fnv.New()
	sum = fnv.Add(sum, "Row, row, row your boat,")
	sum = fnv.Add(sum, "Gently down the stream.")
	fmt.Printf("sum: %d", sum)
	// Output:
	// sum: 13308045931649024798
}
