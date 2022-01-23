package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestTakeWhile(t *testing.T) {
	iter := Ints[int]().TakeWhile(func(i int) bool { return i < 10 })

	// test.Assert(iter.HasNext(), t)
	// test.Assert(iter.HasNext(), t)

	test.AssertDeepEq(
		iter.Collect(),
		Ints[int]().Take(10).Collect(),
		t)

	iter = Ints[int]().Take(0).TakeWhile(func(i int) bool { return i < 10 })

	// test.Assert(!iter.HasNext(), t)

	iter.Collect()

	// test.Assert(!iter.HasNext(), t)

	_, err := iter.Next()

	test.AssertNonNil(err, t)
}

func BenchmarkTakeWhile(b *testing.B) {
	Ints[int]().TakeWhile(func(i int) bool {
		return i < b.N
	}).Consume()
}
