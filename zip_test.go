package iter

import (
	"reflect"
	"testing"
)

// TODO: uncomment this once ZipSame is fixed
// func TestZipSame(t *testing.T) {
// 	iter := Range(0, 9, 1)
//
// 	zipIter := iter.ZipSame(iter)
//
// 	actual := zipIter.Collect()
// 	expected := []Pair[rune, int]{
// 		Pair[rune, int]{0, 1},
// 		Pair[rune, int]{2, 3},
// 		Pair[rune, int]{4, 5},
// 		Pair[rune, int]{6, 7},
// 	}
//
// 	if !reflect.DeepEqual(actual, expected) {
// 		t.Fatalf("got %v, expected %v", actual, expected)
// 	}
// }

func TestZip(t *testing.T) {
	runeIter := Elems([]rune{'a', 'b', 'c', 'd'})
	intIter := Range(1, 1000000000000000, 1)

	zipIter := Zip(runeIter, intIter)

	actual := zipIter.Collect()
	// TODO: file a bug report for this warning on the gopls repo
	expected := []Pair[rune, int]{
		Pair[rune, int]{'a', 1},
		Pair[rune, int]{'b', 2},
		Pair[rune, int]{'c', 3},
		Pair[rune, int]{'d', 4},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
