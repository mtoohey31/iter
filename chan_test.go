package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestMsgs(t *testing.T) {
	ch := make(chan int)
	iter := Msgs(&ch)
	expected := []int{2, 7, 31, 645}

	go func() {
		for _, v := range expected {
			ch <- v
		}
		close(ch)
	}()

	actualStart := iter.Take(2).Collect()

	test.Assert(iter.HasNext(), t)
	test.Assert(iter.HasNext(), t)
	test.AssertDeepEq(append(actualStart, iter.Collect()...), expected, t)

	test.Assert(!iter.HasNext(), t)
}

func BenchmarkMsgs(b *testing.B) {
	ch := make(chan int)

	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	Msgs(&ch).Consume()
}
