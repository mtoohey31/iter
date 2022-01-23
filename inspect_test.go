package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestInspect(t *testing.T) {
	iter := IntsFrom(1).Take(10)

	actualNum := 0
	expectedNumBefore := 0
	expectedNumAfter := 55

	newIter := iter.Inspect(func(n int) {
		actualNum = actualNum + n
	})

	test.AssertEq(actualNum, expectedNumBefore, t)
	// test.Assert(newIter.HasNext(), t)

	actualSlice := newIter.Collect()
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	test.AssertEq(actualNum, expectedNumAfter, t)
	test.AssertDeepEq(actualSlice, expectedSlice, t)
	// test.Assert(!newIter.HasNext(), t)
}

func BenchmarkInspect(b *testing.B) {
	Ints[int]().Inspect(func(i int) {}).Take(b.N).Consume()
}
