package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, expected, append(actualStart, iter.Collect()...))
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

	assert.Equal(t, expected, actual)
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
