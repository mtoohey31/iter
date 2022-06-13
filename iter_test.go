package iter

import (
	"errors"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mp = runtime.GOMAXPROCS(0)

func TestGoConsume(t *testing.T) {
	iter := Ints[int]().Take(10).Mutex()
	iter.GoConsume(mp)

	_, ok := iter()
	assert.False(t, ok)
}

func BenchmarkGoConsume(b *testing.B) {
	Ints[int]().Take(b.N).Inspect(func(int) { time.Sleep(time.Millisecond) }).GoConsume(b.N)
}

func TestCollect(t *testing.T) {
	expected := []string{"item1", "item2"}
	assert.Equal(t, expected, Elems(expected).Collect())
}

func BenchmarkCollect(b *testing.B) {
	Ints[int]().Take(b.N).Collect()
}

func TestGoCollect(t *testing.T) {
	expected := []string{"item1", "item2", "item3", "item4"}
	assert.ElementsMatch(t, expected, Elems(expected).Mutex().GoCollect(mp))
}

func TestCollectInto(t *testing.T) {
	actual := make([]int, 6)
	expected := []int{0, 1, 2, 3, 4, 5}

	assert.Equal(t, 6, Ints[int]().Take(6).CollectInto(actual))
	assert.Equal(t, expected, actual)

	assert.Equal(t, 5, Ints[int]().Take(5).CollectInto(actual))
}

func BenchmarkCollectInto(b *testing.B) {
	slice := make([]int, b.N)
	Ints[int]().CollectInto(slice)
}

func TestGoCollectInto(t *testing.T) {
	actual := make([]int, 6)
	expected := []int{0, 1, 2, 3, 4, 5}

	Ints[int]().Take(6).Mutex().GoCollectInto(actual, mp)
	assert.ElementsMatch(t, expected, actual)

	actual = make([]int, 5)
	Ints[int]().Take(5).Mutex().GoCollectInto(actual, mp)
	assert.ElementsMatch(t, expected[:5], actual)
}

func TestAll(t *testing.T) {
	assert.True(t, !Elems([]int{1, 2}).All(func(i int) bool { return i == 1 }))
	assert.True(t, Elems([]int{1, 2}).All(func(i int) bool { return i != 0 }))
}

func BenchmarkAll(b *testing.B) {
	Ints[int]().Take(b.N).All(func(i int) bool {
		return i >= 0
	})
}

func TestAny(t *testing.T) {
	assert.True(t, Elems([]int{1, 2}).Any(func(i int) bool { return i == 1 }))
	assert.False(t, Elems([]int{1, 2}).Any(func(i int) bool { return i == 0 }))
}

func BenchmarkAny(b *testing.B) {
	Ints[int]().Take(b.N).Any(func(i int) bool {
		return i < 0
	})
}

func TestCount(t *testing.T) {
	assert.Equal(t, 2, Elems([]int{1, 2}).Count())
}

func BenchmarkCount(b *testing.B) {
	Ints[int]().Take(b.N).Count()
}

func TestFind(t *testing.T) {
	actual, _ := Ints[int]().Find(func(i int) bool {
		return i == 7
	})

	assert.Equal(t, 7, actual)

	_, ok := Elems([]bool{}).Find(func(b bool) bool {
		return true
	})

	assert.False(t, ok)
}

func TestFindMapEndo(t *testing.T) {
	actual, _ := Ints[int]().FindMapEndo(func(i int) (int, error) {
		if i == 7 {
			return i, nil
		} else {
			return 0, errors.New("")
		}
	})

	assert.Equal(t, 7, actual)

	_, ok := Elems([]bool{}).FindMapEndo(func(b bool) (bool, error) {
		return true, nil
	})

	assert.False(t, ok)
}

func TestFindMap(t *testing.T) {
	actual, _ := FindMap(Ints[int](), func(i int) (int, error) {
		if i == 7 {
			return i, nil
		} else {
			return 0, errors.New("")
		}
	})

	assert.Equal(t, 7, actual)

	_, ok := FindMap(Elems([]bool{}), func(b bool) (bool, error) {
		return true, nil
	})

	assert.False(t, ok)
}

func TestFoldEndo(t *testing.T) {
	iter := Elems([]string{"quick", "brown", "fox"})

	actual := iter.FoldEndo("the", func(curr string, next string) string {
		return curr + " " + next
	})

	assert.Equal(t, "the quick brown fox", actual)
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

	assert.Equal(t, 16, actual)
}

func BenchmarkFold(b *testing.B) {
	Fold(Ints[int]().Take(b.N), 0, func(p, n int) int {
		return p + n
	})
}

func TestForEach(t *testing.T) {
	actual := 0
	IntsFrom(1).Take(10).ForEach(func(n int) { actual = actual + n })
	assert.Equal(t, 55, actual)
}

func BenchmarkForEach(b *testing.B) {
	Ints[int]().Take(b.N).ForEach(func(i int) {})
}

func TestForEachParallel(t *testing.T) {
	actual := 0
	IntsFrom(1).Take(10).ForEachParallel(func(n int) { actual = actual + n })
	assert.Equal(t, 55, actual)
}

func BenchmarkForEachParallel(b *testing.B) {
	Ints[int]().Take(b.N).ForEachParallel(func(i int) {})
}

func TestLast(t *testing.T) {
	actual, _ := IntsFrom(1).Take(10).Last()

	assert.Equal(t, 10, actual)

	_, ok := Elems([]bool{}).Last()

	assert.False(t, ok)
}

func BenchmarkLast(b *testing.B) {
	Ints[int]().Take(b.N).Last()
}

func TestNth(t *testing.T) {
	actual, _ := IntsFrom(1).Take(10).Nth(7)

	assert.Equal(t, 7, actual)

	_, ok := IntsFrom(1).Take(10).Nth(17)

	assert.False(t, ok)

	_, ok = IntsFrom(1).Take(10).Nth(11)

	assert.False(t, ok)
}

func BenchmarkNth(b *testing.B) {
	Ints[int]().Nth(b.N)
}

func TestTryFoldEndo(t *testing.T) {
	actual, err := IntsBy(2).Take(3).TryFoldEndo(0, func(curr int, next int) (int, error) {
		if next%2 == 0 {
			return curr + next, nil
		} else {
			return 0, errors.New("")
		}
	})

	assert.NoError(t, err)
	assert.Equal(t, 6, actual)

	_, err = Ints[int]().Take(5).TryFoldEndo(0, func(curr int, next int) (int, error) {
		if next%2 == 0 {
			return curr + next, nil
		} else {
			return 0, errors.New("")
		}
	})

	assert.Error(t, err)
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

	assert.NoError(t, err)
	assert.Equal(t, 10, actual)

	_, err = TryFold(Elems([]string{"1", "2", "not a number", "4"}), 0, func(curr int, next string) (int, error) {
		v, err := strconv.Atoi(next)
		if err == nil {
			return curr + v, nil
		} else {
			return 0, err
		}
	})

	assert.Error(t, err)
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

	assert.NoError(t, err)
	assert.Equal(t, 10, actual)

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

	assert.Error(t, err)
	assert.Equal(t, 3, actual)
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

	assert.True(t, ok)
	assert.Equal(t, 4, actual)
}

func BenchmarkReduce(b *testing.B) {
	Ints[int]().Take(b.N).Reduce(func(p, n int) int {
		return 0
	})
}

func TestPosition(t *testing.T) {
	assert.Equal(t, 3, Position(Ints[int](), func(i int) bool { return i == 3 }))
	assert.Equal(t, -1, Position(Ints[int]().Take(0), func(i int) bool { return i == 3 }))
}

func BenchmarkPosition(b *testing.B) {
	Position(Ints[int](), func(i int) bool { return i == b.N })
}

func TestRev(t *testing.T) {
	assert.Equal(t, []int{4, 3, 2, 1, 0}, Ints[int]().Take(5).Rev().Collect())
}
func BenchmarkRev(b *testing.B) {
	Ints[int]().Take(b.N).Rev()
}
