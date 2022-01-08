package iter

import (
	"reflect"
	"testing"
)

func TestCollect(t *testing.T) {
	expected := []string{"item1", "item2"}
	var iter *Iter[string] = Elems(expected)

	actual := iter.Collect()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func TestAll(t *testing.T) {
	var iter *Iter[int] = Elems([]int{1, 2})

	if iter.All(func(i int) bool { return i == 1 }) {
		t.Fatalf("got %v, expected %v", true, false)
	}
}

func TestAny(t *testing.T) {
	var iter *Iter[int] = Elems([]int{1, 2})

	if !iter.Any(func(i int) bool { return i == 1 }) {
		t.Fatalf("got %v, expected %v", false, true)
	}
}

func TestCount(t *testing.T) {
	var iter *Iter[int] = Elems([]int{1, 2})

	actual := iter.Count()
	expected := 2

	if actual != expected {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
