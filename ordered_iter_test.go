package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestMin(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, ok := Min(ordered)

	test.Assert(ok, t)
	test.AssertEq(actual, 0, t)

	_, ok = Min(ordered)

	test.Assert(!ok, t)

	Min(IntsBy(-1).Take(2))
}

func BenchmarkMin(b *testing.B) {
	Min(Ints[int]().Take(b.N))
}

func TestMinByKey(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, ok := MinByKey(ordered, func(n int) int {
		return n * -1
	})

	test.Assert(ok, t)
	test.AssertEq(actual, 9, t)

	_, ok = MinByKey(ordered, func(n int) int {
		return n * -1
	})

	test.Assert(!ok, t)

	ordered = IntsBy(-1).Take(10)

	actual, ok = MinByKey(ordered, func(n int) int {
		return n * -1
	})

	test.Assert(ok, t)
	test.AssertEq(actual, 0, t)
}

func BenchmarkMinByKey(b *testing.B) {
	MinByKey(Ints[int]().Take(b.N), func(n int) int {
		return n
	})
}

func TestMax(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, ok := Max(ordered)

	test.Assert(ok, t)
	test.AssertEq(actual, 9, t)

	_, ok = Max(ordered)

	test.Assert(!ok, t)

	Max(IntsBy(-1).Take(2))
}

func BenchmarkMax(b *testing.B) {
	Max(Ints[int]().Take(b.N))
}

func TestMaxByKey(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, ok := MaxByKey(ordered, func(n int) int {
		return n * -1
	})

	test.Assert(ok, t)
	test.AssertEq(actual, 0, t)

	_, ok = MaxByKey(ordered, func(n int) int {
		return n * -1
	})

	test.Assert(!ok, t)

	ordered = IntsBy(-1).Take(10)

	actual, ok = MaxByKey(ordered, func(n int) int {
		return n * -1
	})

	test.Assert(ok, t)
	test.AssertEq(actual, -9, t)
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
