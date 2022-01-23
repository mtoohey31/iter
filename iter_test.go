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
	Ints[int]().Take(b.N).Collect()
}

func TestCollectInto(t *testing.T) {
	actual := make([]int, 6)
	expected := []int{0, 1, 2, 3, 4, 5}

	test.AssertEq(Ints[int]().Take(6).CollectInto(actual), 6, t)
	test.AssertDeepEq(actual, expected, t)

	test.AssertEq(Ints[int]().Take(5).CollectInto(actual), 5, t)
}

func BenchmarkCollectInto(b *testing.B) {
	slice := make([]int, b.N)
	Ints[int]().CollectInto(slice)
}

func TestAll(t *testing.T) {
	test.Assert(!Elems([]int{1, 2}).All(func(i int) bool { return i == 1 }), t)
	test.Assert(Elems([]int{1, 2}).All(func(i int) bool { return i != 0 }), t)
}

func BenchmarkAll(b *testing.B) {
	Ints[int]().Take(b.N).All(func(i int) bool {
		return i >= 0
	})
}

func TestAny(t *testing.T) {
	test.Assert(Elems([]int{1, 2}).Any(func(i int) bool { return i == 1 }), t)
	test.Assert(!Elems([]int{1, 2}).Any(func(i int) bool { return i == 0 }), t)
}

func BenchmarkAny(b *testing.B) {
	Ints[int]().Take(b.N).Any(func(i int) bool {
		return i < 0
	})
}

func TestCount(t *testing.T) {
	test.AssertEq(Elems([]int{1, 2}).Count(), 2, t)
}

func BenchmarkCount(b *testing.B) {
	Ints[int]().Take(b.N).Count()
}

func TestFind(t *testing.T) {
	actual, _ := Ints[int]().Find(func(i int) bool {
		return i == 7
	})

	test.AssertEq(actual, 7, t)

	_, ok := Elems([]bool{}).Find(func(b bool) bool {
		return true
	})

	test.AssertNonNil(ok, t)
}

func TestFindMapEndo(t *testing.T) {
	actual, _ := Ints[int]().FindMapEndo(func(i int) (int, error) {
		if i == 7 {
			return i, nil
		} else {
			return 0, errors.New("")
		}
	})

	test.AssertEq(actual, 7, t)

	_, ok := Elems([]bool{}).FindMapEndo(func(b bool) (bool, error) {
		return true, nil
	})

	test.AssertNonNil(ok, t)
}

func TestFindMap(t *testing.T) {
	actual, _ := FindMap(Ints[int](), func(i int) (int, error) {
		if i == 7 {
			return i, nil
		} else {
			return 0, errors.New("")
		}
	})

	test.AssertEq(actual, 7, t)

	_, ok := FindMap(Elems([]bool{}), func(b bool) (bool, error) {
		return true, nil
	})

	test.AssertNonNil(ok, t)
}

func TestFoldEndo(t *testing.T) {
	iter := Elems([]string{"quick", "brown", "fox"})

	actual := iter.FoldEndo("the", func(curr string, next string) string {
		return curr + " " + next
	})

	test.AssertEq(actual, "the quick brown fox", t)
}

func BenchmarkFoldEndo(b *testing.B) {
	Ints[int]().Take(b.N).FoldEndo(0, func(p, n int) int {
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
	Fold(Ints[int]().Take(b.N), 0, func(p, n int) int {
		return p + n
	})
}

func TestForEach(t *testing.T) {
	actual := 0
	IntsFrom(1).Take(10).ForEach(func(n int) { actual = actual + n })
	test.AssertEq(actual, 55, t)
}

func BenchmarkForEach(b *testing.B) {
	Ints[int]().Take(b.N).ForEach(func(i int) {})
}

func TestLast(t *testing.T) {
	actual, _ := IntsFrom(1).Take(10).Last()

	test.AssertEq(actual, 10, t)

	_, ok := Elems([]bool{}).Last()

	test.AssertNonNil(ok, t)
}

func BenchmarkLast(b *testing.B) {
	Ints[int]().Take(b.N).Last()
}

func TestNth(t *testing.T) {
	actual, _ := IntsFrom(1).Take(10).Nth(7)

	test.AssertEq(actual, 7, t)

	_, ok := IntsFrom(1).Take(10).Nth(17)

	test.AssertNonNil(ok, t)

	_, ok = IntsFrom(1).Take(10).Nth(11)

	test.AssertNonNil(ok, t)
}

func BenchmarkNth(b *testing.B) {
	Ints[int]().Nth(b.N)
}

func TestPartition(t *testing.T) {
	actualA, actualB := Ints[int]().Take(4).Partition(func(i int) bool { return i%2 == 0 })

	test.AssertDeepEq(actualA, []int{0, 2}, t)
	test.AssertDeepEq(actualB, []int{1, 3}, t)
}

func BenchmarkPartition(b *testing.B) {
	Ints[int]().Take(b.N).Partition(func(i int) bool {
		return i%2 == 0
	})
}

func TestTryFoldEndo(t *testing.T) {
	actual, ok := IntsBy(2).Take(3).TryFoldEndo(0, func(curr int, next int) (int, error) {
		if next%2 == 0 {
			return curr + next, nil
		} else {
			return 0, errors.New("")
		}
	})

	test.AssertNil(ok, t)
	test.AssertEq(actual, 6, t)

	_, ok = Ints[int]().Take(5).TryFoldEndo(0, func(curr int, next int) (int, error) {
		if next%2 == 0 {
			return curr + next, nil
		} else {
			return 0, errors.New("")
		}
	})

	test.AssertNonNil(ok, t)
}

func BenchmarkTryFoldEndo(b *testing.B) {
	Ints[int]().Take(b.N).TryFoldEndo(0, func(curr, next int) (int, error) {
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
	TryFold(Ints[int]().Take(b.N), 0, func(curr, next int) (int, error) {
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
	Ints[int]().Take(b.N).TryForEach(func(i int) error { return nil })
}

func TestReduce(t *testing.T) {
	actual, ok := Ints[int]().Take(5).Reduce(func(curr int, next int) int {
		if next > curr {
			return next
		} else {
			return curr
		}
	})

	test.Assert(ok, t)
	test.AssertEq(actual, 4, t)
}

func BenchmarkReduce(b *testing.B) {
	Ints[int]().Take(b.N).Reduce(func(p, n int) int {
		return 0
	})
}

func TestPosition(t *testing.T) {
	test.AssertEq(Position(Ints[int](), func(i int) bool { return i == 3 }), 3, t)
	test.AssertEq(Position(Ints[int]().Take(0), func(i int) bool { return i == 3 }), -1, t)
}

func BenchmarkPosition(b *testing.B) {
	Position(Ints[int](), func(i int) bool { return i == b.N })
}

func TestRev(t *testing.T) {
	test.AssertDeepEq(Ints[int]().Take(5).Rev().Collect(), []int{4, 3, 2, 1, 0}, t)
}
func BenchmarkRev(b *testing.B) {
	Ints[int]().Take(b.N).Rev()
}
