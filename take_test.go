package iter

import (
	"mtoohey.com/iter/test"
	"testing"
)

func TestTake(t *testing.T) {
	iter := Ints[int]().Take(10)
	test.AssertEq(iter.Take(5).Count(), 5, t)
	test.AssertEq(iter.Count(), 5, t)
}

func BenchmarkTake(b *testing.B) {
	Ints[int]().Take(b.N).Consume()
}
