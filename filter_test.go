package iter

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	var iter Iter[int] = Filter(FromSlice([]int{1, 2, 3, 4}), func(i int) bool { return i%2 == 0 })

	actual := Collect(iter)
	expected := []int{2, 4}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
