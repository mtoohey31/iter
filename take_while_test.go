package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestTakeWhile(t *testing.T) {
	test.AssertDeepEq(
		Ints[int]().TakeWhile(func(i int) bool { return i < 10 }).Collect(),
		Ints[int]().Take(10).Collect(),
		t)
}

func BenchmarkTakeWhile(b *testing.B) {
	Ints[int]().TakeWhile(func(i int) bool {
		return i < b.N
	}).Consume()
}
