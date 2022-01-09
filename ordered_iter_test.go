package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestMin(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, err := Min(ordered)

	test.AssertNil(err, t)
	test.AssertEq(actual, 0, t)

	_, err = Min(ordered)

	test.AssertNonNil(err, t)
}

func BenchmarkMin(b *testing.B) {
	Min(Ints[int]().Take(b.N))
}

func TestMinByKey(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, err := MinByKey(ordered, func(n int) int {
		return n * -1
	})

	test.AssertNil(err, t)
	test.AssertEq(actual, 9, t)

	_, err = MinByKey(ordered, func(n int) int {
		return n * -1
	})

	test.AssertNonNil(err, t)
}

func BenchmarkMinByKey(b *testing.B) {
	MinByKey(Ints[int]().Take(b.N), func(n int) int {
		return n
	})
}

func TestMax(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, err := Max(ordered)

	test.AssertNil(err, t)
	test.AssertEq(actual, 9, t)

	_, err = Max(ordered)

	test.AssertNonNil(err, t)
}

func BenchmarkMax(b *testing.B) {
	Max(Ints[int]().Take(b.N))
}

func TestMaxByKey(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, err := MaxByKey(ordered, func(n int) int {
		return n * -1
	})

	test.AssertNil(err, t)
	test.AssertEq(actual, 0, t)

	_, err = MaxByKey(ordered, func(n int) int {
		return n * -1
	})

	test.AssertNonNil(err, t)
}

func BenchmarkMaxByKey(b *testing.B) {
	MaxByKey(Ints[int]().Take(b.N), func(n int) int {
		return n
	})
}

func TestSum(t *testing.T) {
	test.AssertEq(Sum(Ints[int]().Take(10)), 45, t)
}

func BenchmarkSum(b *testing.B) {
	Sum(Ints[int]().Take(b.N))
}
