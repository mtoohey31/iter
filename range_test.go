package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestRangeIter(t *testing.T) {
	test.AssertDeepEq(Range(1, 7, 2).Collect(), []int{1, 3, 5}, t)
}

func TestInfRangeIter(t *testing.T) {
	test.AssertDeepEq(InfRange(7, -2).Take(7).Collect(),
		[]int{7, 5, 3, 1, -1, -3, -5}, t)
}
