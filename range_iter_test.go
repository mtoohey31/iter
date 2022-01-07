package iter

import (
	"reflect"
	"testing"
)

func TestRangeIter(t *testing.T) {
	expected := []int{1, 3, 5}
	var iter Iter[int] = Range(1, 7, 2)

	actual := Collect(iter)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
