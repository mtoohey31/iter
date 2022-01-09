package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestMin(t *testing.T) {
	ordered := OrderedIter[int](*Ints[int]().Take(10))

	actual, err := ordered.Min()

	test.AssertNil(err, t)
	test.AssertEq(actual, 0, t)

	_, err = ordered.Min()

	test.AssertNonNil(err, t)
}

func BenchmarkMin(b *testing.B) {
	ordered := OrderedIter[int](*Ints[int]().Take(b.N))
	ordered.Min()
}

func TestMinByKey(t *testing.T) {
	ordered := OrderedIter[int](*Ints[int]().Take(10))

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

func BenchmarkMinByKey(b *testing.B) {
	ordered := OrderedIter[int](*Ints[int]().Take(b.N))
	MinByKey(&ordered, func(n int) int {
		return n
	})
}

func TestMax(t *testing.T) {
	ordered := OrderedIter[int](*Ints[int]().Take(10))

	actual, err := ordered.Max()

	test.AssertNil(err, t)
	test.AssertEq(actual, 9, t)

	_, err = ordered.Max()

	test.AssertNonNil(err, t)
}

func BenchmarkMax(b *testing.B) {
	ordered := OrderedIter[int](*Ints[int]().Take(b.N))
	ordered.Max()
}

func TestMaxByKey(t *testing.T) {
	ordered := OrderedIter[int](*Ints[int]().Take(10))

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

func BenchmarkMaxByKey(b *testing.B) {
	ordered := OrderedIter[int](*Ints[int]().Take(b.N))
	MaxByKey(&ordered, func(n int) int {
		return n
	})
}

func TestSum(t *testing.T) {
	ordered := OrderedIter[int](*Ints[int]().Take(10))
	test.AssertEq(ordered.Sum(), 45, t)
}

func BenchmarkSum(b *testing.B) {
	ordered := OrderedIter[int](*Ints[int]().Take(b.N))
	ordered.Sum()
}
