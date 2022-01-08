package iter

import (
	"github.com/barweiss/go-tuple"

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
// 	expected := []tuple.T2[int, int]{
// 		tuple.New2(0, 1),
// 		tuple.New2(2, 3),
// 		tuple.New2(4, 5),
// 		tuple.New2(6, 7),
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
	expected := []tuple.T2[rune, int]{
		tuple.New2('a', 1),
		tuple.New2('b', 2),
		tuple.New2('c', 3),
		tuple.New2('d', 4),
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
