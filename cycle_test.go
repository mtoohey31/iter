package iter

import (
	"reflect"
	"testing"
)

func TestCycle(t *testing.T) {
	iter := Elems([]int{1, 2}).Cycle()

	var actual []int

	for i := 0; i < 6; i++ {
		next, _ := iter.Next()
		actual = append(actual, next)
	}

	expected := []int{1, 2, 1, 2, 1, 2}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
