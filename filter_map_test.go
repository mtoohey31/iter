package iter

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestFilterMapSame(t *testing.T) {
	var iter *Iter[int] = Elems([]int{1, 2, 3, 4}).FilterMapSame(func(i int) (int, error) {
		if i%2 != 0 {
			return 0, errors.New("")
		} else {
			return i * 2, nil
		}
	})

	actual := iter.Collect()
	expected := []int{4, 8}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func TestFilterMap(t *testing.T) {
	var iter *Iter[int] = FilterMap(Elems([]string{"1", "nope", "2", "un-uh"}), func(s string) (int, error) {
		return strconv.Atoi(s)
	})

	actual := iter.Collect()
	expected := []int{1, 2}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
