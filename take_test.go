package iter

import (
	"mtoohey.com/iter/test"
	"testing"
)

func TestTake(t *testing.T) {
	iter := Range(0, 10, 1)
	test.AssertEq(iter.Take(5).Count(), 5, t)
	test.AssertEq(iter.Count(), 5, t)
}

func BenchmarkTake(b *testing.B) {
	InfRange(0, 1).Take(b.N).Consume()
}
