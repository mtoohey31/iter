package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestInspect(t *testing.T) {
	iter := Range(1, 11, 1)

	actualNum := 0
	expectedNumBefore := 0
	expectedNumAfter := 55

	newIter := iter.Inspect(func(n int) {
		actualNum = actualNum + n
	})

	test.AssertEq(actualNum, expectedNumBefore, t)

	actualSlice := newIter.Collect()
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	test.AssertEq(actualNum, expectedNumAfter, t)

	test.AssertDeepEq(actualSlice, expectedSlice, t)
}

func BenchmarkInspect(b *testing.B) {
	InfRange(0, 1).Inspect(func(i int) {}).Take(b.N).Consume()
}
