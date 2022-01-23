package iter

import (
	"errors"
	"testing"

	"mtoohey.com/iter/test"
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

	// test.Assert(iter.HasNext(), t)
	// test.Assert(iter.HasNext(), t)
	test.AssertDeepEq(iter.Take(5).Collect(), []int{0, 0, 0, 0, 0}, t)

	b = true

	// test.Assert(!iter.HasNext(), t)
	// test.Assert(!iter.HasNext(), t)

	_, err := iter()

	test.AssertNonNil(err, t)

	iter = GenWhile(func() (int, error) {
		if b {
			return 0, errors.New("")
		} else {
			return 0, nil
		}
	})

	_, err = iter()

	test.AssertNonNil(err, t)
}

func BenchmarkGenWhile(b *testing.B) {
	GenWhile(func() (int, error) {
		return 0, nil
	}).Take(b.N).Consume()
}
