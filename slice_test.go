package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestElems(t *testing.T) {
	expected := []string{"item1", "item2"}
	test.AssertDeepEq(Elems(expected).Collect(), expected, t)
}

func BenchmarkElems(b *testing.B) {
	slice := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		slice[i] = i
	}

	Elems(slice)
}
