package iter

import (
	"reflect"
	"strings"
	"testing"
)

func TestFlatMapSame(t *testing.T) {
	initial := []int{1, 2, 3}
	iter := Elems(initial).FlatMapSame(func(i int) *Iter[int] {
		return Range(i, i+2, 1)
	})

	actual := iter.Collect()
	expected := []int{1, 2, 2, 3, 3, 4}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func TestFlatMap(t *testing.T) {
	initial := []string{"alpha", "beta", "gamma"}
	iter := FlatMap(Elems(initial), func(s string) *Iter[rune] {
		return Runes(s)
	})

	actual := string(iter.Collect())
	expected := strings.Join(initial, "")

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
