package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestMsgs(t *testing.T) {
	ch := make(chan int)
	expected := []int{2, 7, 31, 645}

	go func() {
		for _, v := range expected {
			ch <- v
		}
		close(ch)
	}()

	test.AssertDeepEq(Msgs(&ch).Collect(), expected, t)
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
