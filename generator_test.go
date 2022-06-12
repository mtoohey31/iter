package iter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenWhile(t *testing.T) {
	b := false

	iter := GenWhile(func() (int, error) {
		if b {
			return 0, errors.New("")
		} else {
			return 0, nil
		}
	})

	assert.Equal(t, []int{0, 0, 0, 0, 0}, iter.Take(5).Collect())

	b = true

	_, ok := iter()

	assert.False(t, ok)

	iter = GenWhile(func() (int, error) {
		if b {
			return 0, errors.New("")
		} else {
			return 0, nil
		}
	})

	_, ok = iter()

	assert.False(t, ok)
}

func BenchmarkGenWhile(b *testing.B) {
	GenWhile(func() (int, error) {
		return 0, nil
	}).Take(b.N).Consume()
}
