package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestReceive(t *testing.T) {
	ch := make(chan int)
	iter := Receive(&ch)
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

func BenchmarkReceive(b *testing.B) {
	ch := make(chan int)

	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	Receive(&ch).Consume()
}

func TestSend(t *testing.T) {
	ch := make(chan int)
	expected := []int{5, 8, 11, 14}

	go func() {
		IntsFromBy(5, 3).Take(4).Send(&ch)
		close(ch)
	}()

	actual := make([]int, 4)
	i := 0
	for v := range ch {
		actual[i] = v
		i++
	}

	test.AssertDeepEq(actual, expected, t)
}

func BenchmarkSend(b *testing.B) {
	ch := make(chan int)

	go func() {
		Ints[int]().Take(b.N).Send(&ch)
		close(ch)
	}()

	for range ch {
	}
}
