package iter

import (
	"reflect"
	"testing"
)

func TestChain(t *testing.T) {
	iter1 := Elems([]int{1, 2})
	iter2 := Elems([]int{3, 4})

	actual := iter1.Chain(iter2).Collect()
	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
