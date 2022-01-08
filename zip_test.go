package iter

import (
	"github.com/barweiss/go-tuple"

	"mtoohey.com/iter/test"

	"testing"
)

// TODO: uncomment this once ZipSame is fixed
// func TestZipSame(t *testing.T) {
// 	iter := Range(0, 9, 1)
//
// 	zipIter := iter.ZipSame(iter)
//
// 	expected := []tuple.T2[int, int]{
// 		tuple.New2(0, 1),
// 		tuple.New2(2, 3),
// 		tuple.New2(4, 5),
// 		tuple.New2(6, 7),
// 	}
//
// 	test.AssertDeepEq(zipIter.Collect(), expected, t)
// }

func TestZip(t *testing.T) {
	iter := Zip(Elems([]rune{'a', 'b', 'c', 'd'}), InfRange(1, 1))

	expected := []tuple.T2[rune, int]{
		tuple.New2('a', 1),
		tuple.New2('b', 2),
		tuple.New2('c', 3),
		tuple.New2('d', 4),
	}

	test.AssertDeepEq(iter.Collect(), expected, t)
}
