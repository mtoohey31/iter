package iter

import (
	"reflect"
	"testing"
)

func TestCollect(t *testing.T) {
	expected := []string{"item1", "item2"}
	iter := Elems(expected)

	actual := iter.Collect()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func TestAll(t *testing.T) {
	iter := Elems([]int{1, 2})

	if iter.All(func(i int) bool { return i == 1 }) {
		t.Fatalf("got %v, expected %v", true, false)
	}
}

func TestAny(t *testing.T) {
	iter := Elems([]int{1, 2})

	if !iter.Any(func(i int) bool { return i == 1 }) {
		t.Fatalf("got %v, expected %v", false, true)
	}
}

func TestCount(t *testing.T) {
	iter := Elems([]int{1, 2})

	actual := iter.Count()
	expected := 2

	if actual != expected {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func TestFoldSame(t *testing.T) {
	iter := Elems([]string{"quick", "brown", "fox"})

	actual := iter.FoldSame("the", func(curr string, next string) string {
		return curr + " " + next
	})
	expected := "the quick brown fox"

	if actual != expected {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func TestFold(t *testing.T) {
	iter := Elems([]string{"the", "quick", "brown", "fox"})

	actual := Fold(iter, 0, func(curr int, next string) int {
		return curr + len(next)
	})
	expected := 16

	if actual != expected {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func TestForEach(t *testing.T) {
	iter := Range(1, 11, 1)

	actual := 0
	expected := 55

	iter.ForEach(func(n int) { actual = actual + n })

	if actual != expected {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
