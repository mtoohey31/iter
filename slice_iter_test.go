package iter

import (
	"reflect"
	"testing"
)

func TestArrIter(t *testing.T) {
	expected := []string{"item1", "item2"}
	var iter Iter[string] = FromSlice(expected)

	actual := Collect(iter)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
