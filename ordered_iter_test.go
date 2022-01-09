package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestMin(t *testing.T) {
	ordered := OrderedIter[int](*Range(0, 10, 1))

	actual, err := ordered.Min()

	test.AssertNil(err, t)
	test.AssertEq(actual, 0, t)

	_, err = ordered.Min()

	test.AssertNonNil(err, t)
}

func TestMinByKey(t *testing.T) {
	ordered := OrderedIter[int](*Range(0, 10, 1))

	actual, err := MinByKey(&ordered, func(n int) int {
		return n * -1
	})

	test.AssertNil(err, t)
	test.AssertEq(actual, 9, t)

	_, err = MinByKey(&ordered, func(n int) int {
		return n * -1
	})

	test.AssertNonNil(err, t)
}

func TestMax(t *testing.T) {
	ordered := OrderedIter[int](*Range(0, 10, 1))

	actual, err := ordered.Max()

	test.AssertNil(err, t)
	test.AssertEq(actual, 9, t)

	_, err = ordered.Max()

	test.AssertNonNil(err, t)
}

func TestMaxByKey(t *testing.T) {
	ordered := OrderedIter[int](*Range(0, 10, 1))

	actual, err := MaxByKey(&ordered, func(n int) int {
		return n * -1
	})

	test.AssertNil(err, t)
	test.AssertEq(actual, 0, t)

	_, err = MaxByKey(&ordered, func(n int) int {
		return n * -1
	})

	test.AssertNonNil(err, t)
}

func TestSum(t *testing.T) {
	ordered := OrderedIter[int](*Range(0, 10, 1))
	test.AssertEq(ordered.Sum(), 45, t)
}
