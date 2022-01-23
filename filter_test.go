package iter

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	iter := Elems([]int{1, 2, 3, 4}).Filter(func(i int) bool { return i%2 == 0 })

	actualFirst, _ := iter()
	expected := []int{2, 4}

	assert.Equal(t, append([]int{actualFirst}, iter.Collect()...), expected)
}

func BenchmarkFilter(b *testing.B) {
	Ints[int]().Filter(func(i int) bool {
		return i%2 == 0
	}).Take(b.N).Consume()
}

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

	assert.Equal(t, actual, expected)
}

func TestFilterMap(t *testing.T) {
	iter := FilterMap(Elems([]string{"1", "nope", "2", "un-uh"}), func(s string) (int, error) {
		return strconv.Atoi(s)
	})

	actualFirst, _ := iter()
	expected := []int{1, 2}

	assert.Equal(t, append([]int{actualFirst}, iter.Collect()...), expected)
}

func BenchmarkFilterMapEndo(b *testing.B) {
	var dummyErr error

	Ints[int]().FilterMapEndo(func(i int) (int, error) {
		if i%2 == 0 {
			return i * 2, nil
		} else {
			return 0, dummyErr
		}
	}).Take(b.N).Consume()
}

func BenchmarkFilterMap(b *testing.B) {
	var dummyErr error

	FilterMap(Ints[int](), func(i int) (int, error) {
		if i%2 == 0 {
			return i * 2, nil
		} else {
			return 0, dummyErr
		}
	}).Take(b.N).Consume()
}
