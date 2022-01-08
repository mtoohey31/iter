package iter

import (
	"reflect"
	"testing"
)

func TestRangeIter(t *testing.T) {
	iter := Range(1, 7, 2)

	actual := iter.Collect()
	expected := []int{1, 3, 5}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func TestInfRangeIter(t *testing.T) {
	iter := InfRange(7, -2)

	actual := iter.Take(7)
	expected := []int{7, 5, 3, 1, -1, -3, -5}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
