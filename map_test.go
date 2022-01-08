package iter

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	expected := []int{5, 5}
	var iter *Iter[int] = Map[string, int](FromSlice([]string{"item1", "item2"}), func(s string) int { return len(s) })

	actual := iter.Collect()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}

	if len(iter.Collect()) != 0 {
		t.Fatalf("original iterator was not drained")
	}
}
