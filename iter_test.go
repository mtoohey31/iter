package iter

import (
	"errors"
	"strconv"
	"testing"

	"mtoohey.com/iter/test"
)

func TestCollect(t *testing.T) {
	expected := []string{"item1", "item2"}
	test.AssertDeepEq(Elems(expected).Collect(), expected, t)
}

func BenchmarkCollect(b *testing.B) {
	InfRange(0, 1).Take(b.N).Collect()
}

func TestAll(t *testing.T) {
	test.Assert(!Elems([]int{1, 2}).All(func(i int) bool { return i == 1 }), t)
}

func BenchmarkAll(b *testing.B) {
	InfRange(0, 1).Take(b.N).All(func(i int) bool {
		return i >= 0
	})
}

func TestAny(t *testing.T) {
	test.Assert(Elems([]int{1, 2}).Any(func(i int) bool { return i == 1 }), t)
}

func BenchmarkAny(b *testing.B) {
	InfRange(0, 1).Take(b.N).Any(func(i int) bool {
		return i < 0
	})
}

func TestCount(t *testing.T) {
	test.AssertEq(Elems([]int{1, 2}).Count(), 2, t)
}

func BenchmarkCount(b *testing.B) {
	InfRange(0, 1).Take(b.N).Count()
}

func TestFoldEndo(t *testing.T) {
	iter := Elems([]string{"quick", "brown", "fox"})

	actual := iter.FoldEndo("the", func(curr string, next string) string {
		return curr + " " + next
	})

	test.AssertEq(actual, "the quick brown fox", t)
}

func BenchmarkFoldEndo(b *testing.B) {
	InfRange(0, 1).Take(b.N).FoldEndo(0, func(p, n int) int {
		return p + n
	})
}

func TestFold(t *testing.T) {
	iter := Elems([]string{"the", "quick", "brown", "fox"})

	actual := Fold(iter, 0, func(curr int, next string) int {
		return curr + len(next)
	})

	test.AssertEq(actual, 16, t)
}

func BenchmarkFold(b *testing.B) {
	Fold(InfRange(0, 1).Take(b.N), 0, func(p, n int) int {
		return p + n
	})
}

func TestForEach(t *testing.T) {
	actual := 0
	Range(1, 11, 1).ForEach(func(n int) { actual = actual + n })
	test.AssertEq(actual, 55, t)
}

func BenchmarkForEach(b *testing.B) {
	InfRange(0, 1).Take(b.N).ForEach(func(i int) {})
}

func TestLast(t *testing.T) {
	actual, _ := Range(1, 11, 1).Last()
	test.AssertEq(actual, 10, t)
}

func BenchmarkLast(b *testing.B) {
	InfRange(0, 1).Take(b.N).Last()
}

func TestNth(t *testing.T) {
	actual, _ := Range(1, 11, 1).Nth(7)
	test.AssertEq(actual, 7, t)
}

func BenchmarkNth(b *testing.B) {
	InfRange(0, 1).Nth(b.N)
}

func TestPartition(t *testing.T) {
	actualA, actualB := Range(0, 4, 1).Partition(func(i int) bool { return i%2 == 0 })

	test.AssertDeepEq(actualA, []int{0, 2}, t)
	test.AssertDeepEq(actualB, []int{1, 3}, t)
}

func BenchmarkPartition(b *testing.B) {
	InfRange(0, 1).Take(b.N).Partition(func(i int) bool {
		return i%2 == 0
	})
}

func TestTryFoldEndo(t *testing.T) {
	actual, err := Range(0, 5, 2).TryFoldEndo(0, func(curr int, next int) (int, error) {
		if next%2 == 0 {
			return curr + next, nil
		} else {
			return 0, errors.New("")
		}
	})

	test.AssertNil(err, t)
	test.AssertEq(actual, 6, t)

	_, err = Range(0, 5, 1).TryFoldEndo(0, func(curr int, next int) (int, error) {
		if next%2 == 0 {
			return curr + next, nil
		} else {
			return 0, errors.New("")
		}
	})

	test.AssertNonNil(err, t)
}

func BenchmarkTryFoldEndo(b *testing.B) {
	InfRange(0, 1).Take(b.N).TryFoldEndo(0, func(curr, next int) (int, error) {
		return 0, nil
	})
}

func TestTryFold(t *testing.T) {
	actual, err := TryFold(Elems([]string{"1", "2", "3", "4"}), 0, func(curr int, next string) (int, error) {
		v, err := strconv.Atoi(next)
		if err == nil {
			return curr + v, nil
		} else {
			return 0, err
		}
	})

	test.AssertNil(err, t)
	test.AssertEq(actual, 10, t)

	_, err = TryFold(Elems([]string{"1", "2", "not a number", "4"}), 0, func(curr int, next string) (int, error) {
		v, err := strconv.Atoi(next)
		if err == nil {
			return curr + v, nil
		} else {
			return 0, err
		}
	})

	test.AssertNonNil(err, t)
}

func BenchmarkTryFold(b *testing.B) {
	TryFold(InfRange(0, 1).Take(b.N), 0, func(curr, next int) (int, error) {
		return 0, nil
	})
}

func TestTryForEach(t *testing.T) {
	actual := 0
	err := Elems([]string{"1", "2", "3", "4"}).TryForEach(func(s string) error {
		v, err := strconv.Atoi(s)
		if err == nil {
			actual += v
			return nil
		} else {
			return err
		}
	})

	test.AssertNil(err, t)
	test.AssertEq(actual, 10, t)

	actual = 0
	err = Elems([]string{"1", "2", "not a number", "4"}).TryForEach(func(s string) error {
		v, err := strconv.Atoi(s)
		if err == nil {
			actual += v
			return nil
		} else {
			return err
		}
	})

	test.AssertNonNil(err, t)
	test.AssertEq(actual, 3, t)
}

func BenchmarkTryForEach(b *testing.B) {
	InfRange(0, 1).Take(b.N).TryForEach(func(i int) error { return nil })
}

func TestReduce(t *testing.T) {
	actual, err := Range(0, 5, 1).Reduce(func(curr int, next int) int {
		if next > curr {
			return next
		} else {
			return curr
		}
	})

	test.AssertNil(err, t)
	test.AssertEq(actual, 4, t)
}

func BenchmarkReduce(b *testing.B) {
	InfRange(0, 1).Take(b.N).Reduce(func(p, n int) int {
		return 0
	})
}
