package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v3/internal/testutils"
)

func FuzzReceive(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		ch := make(chan byte)

		go func() {
			for _, v := range b {
				ch <- v
			}
			close(ch)
		}()

		assert.Equal(t, b, Receive(&ch).Collect())
	})
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

func FuzzIter_Send(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		ch := make(chan byte)

		go func() {
			Elems(b).Send(&ch)
			close(ch)
		}()

		actual := []byte{}
		for v := range ch {
			actual = append(actual, v)
		}

		assert.Equal(t, b, actual)
	})
}

func BenchmarkIter_Send(b *testing.B) {
	ch := make(chan int)

	go func() {
		Ints[int]().Take(uint(b.N)).Send(&ch)
		close(ch)
	}()

	for range ch {
	}
}
