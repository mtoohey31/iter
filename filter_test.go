package iter

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	var iter *Iter[int] = FromSlice([]int{1, 2, 3, 4}).Filter(func(i int) bool { return i%2 == 0 })

	actual := iter.Collect()
	expected := []int{2, 4}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
