package iter

import (
	"github.com/barweiss/go-tuple"

	"mtoohey.com/iter/test"

	"testing"
)

// TODO: uncomment this once ZipEndo is fixed
// func TestZipEndo(t *testing.T) {
// 	iter := Range(0, 9, 1)
//
// 	zipIter := iter.ZipEndo(iter)
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

func BenchmarkZip(b *testing.B) {
	Zip(InfRange(0, 1), InfRange(0, 1)).Take(b.N).Consume()
}

func TestEnumerate(t *testing.T) {
	expected := []tuple.T2[int, int]{
		tuple.New2(0, 7),
		tuple.New2(1, 5),
		tuple.New2(2, 3),
		tuple.New2(3, 1),
	}

	test.AssertDeepEq(Enumerate(Range(7, 0, -2)).Collect(), expected, t)
}

func BenchmarkEnumerate(b *testing.B) {
	Enumerate(InfRange(0, 1)).Take(b.N).Consume()
}

func TestUnzip(t *testing.T) {
	expected := tuple.New2(Range(0, 10, 1).Collect(), Range(10, 0, -1).Collect())
	v1, v2 := Unzip(Zip(Elems(expected.V1), Elems(expected.V2)))
	test.AssertDeepEq(
		tuple.New2(v1.Collect(), v2.Collect()),
		expected,
		t)
}

func BenchmarkUnzip(b *testing.B) {
	v1, v2 := Unzip(Zip(InfRange(0, 1), InfRange(0, 1)))
	v1.Take(b.N).Consume()
	v2.Take(b.N).Consume()
}
