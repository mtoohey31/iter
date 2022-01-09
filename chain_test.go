package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestChain(t *testing.T) {
	actual := Elems([]int{1, 2}).Chain(Elems([]int{3, 4})).Collect()
	test.AssertDeepEq(actual, []int{1, 2, 3, 4}, t)
}
