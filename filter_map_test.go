package iter

import (
	"errors"
	"strconv"
	"testing"

	"mtoohey.com/iter/test"
)

func TestFilterMapEndo(t *testing.T) {
	iter := Elems([]int{1, 2, 3, 4}).FilterMapEndo(func(i int) (int, error) {
		if i%2 != 0 {
			return 0, errors.New("")
		} else {
			return i * 2, nil
		}
	})

	actual := iter.Collect()
	expected := []int{4, 8}

	test.AssertDeepEq(actual, expected, t)
}

func TestFilterMap(t *testing.T) {
	iter := FilterMap(Elems([]string{"1", "nope", "2", "un-uh"}), func(s string) (int, error) {
		return strconv.Atoi(s)
	})

	actual := iter.Collect()
	expected := []int{1, 2}

	test.AssertDeepEq(actual, expected, t)
}

func BenchmarkFilterMapEndo(b *testing.B) {
	var dummyErr error

	InfRange(0, 1).FilterMapEndo(func(i int) (int, error) {
		if i%2 == 0 {
			return i * 2, nil
		} else {
			return 0, dummyErr
		}
	}).Take(b.N).Consume()
}

func BenchmarkFilterMap(b *testing.B) {
	var dummyErr error

	FilterMap(InfRange(0, 1), func(i int) (int, error) {
		if i%2 == 0 {
			return i * 2, nil
		} else {
			return 0, dummyErr
		}
	}).Take(b.N).Consume()
}
