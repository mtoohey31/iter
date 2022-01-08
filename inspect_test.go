package iter

import (
	"reflect"
	"testing"
)

func TestInspect(t *testing.T) {
	iter := Range(1, 11, 1)

	actualNum := 0
	expectedNumBefore := 0
	expectedNumAfter := 55

	newIter := iter.Inspect(func(n int) {
		actualNum = actualNum + n
	})

	if actualNum != expectedNumBefore {
		t.Fatalf("got %v, expected %v", actualNum, expectedNumBefore)
	}

	actualSlice := newIter.Collect()
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if actualNum != expectedNumAfter {
		t.Fatalf("got %v, expected %v", actualNum, expectedNumAfter)
	}

	if !reflect.DeepEqual(actualSlice, expectedSlice) {
		t.Fatalf("got %v, expected %v", actualSlice, expectedSlice)
	}
}
