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

	test.AssertDeepEq(Msgs(ch).Collect(), expected, t)
}
