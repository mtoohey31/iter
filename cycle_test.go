package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestCycle(t *testing.T) {
	iter := Elems([]int{1, 2}).Cycle()

	var actual []int

	for i := 0; i < 6; i++ {
		next, _ := iter.Next()
		actual = append(actual, next)
	}

	expected := []int{1, 2, 1, 2, 1, 2}

	test.AssertDeepEq(actual, expected, t)
}
