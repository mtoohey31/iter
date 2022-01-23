package iter

import (
	"github.com/barweiss/go-tuple"

	"mtoohey.com/iter/test"

	"testing"
)

func TestZip(t *testing.T) {
	iter := Zip(Elems([]rune{'a', 'b', 'c', 'd'}), IntsFrom(1))

	// test.Assert(iter.HasNext(), t)

	expected := []tuple.T2[rune, int]{
		tuple.New2('a', 1),
		tuple.New2('b', 2),
		tuple.New2('c', 3),
		tuple.New2('d', 4),
	}

	test.AssertDeepEq(iter.Collect(), expected, t)
	// test.Assert(!iter.HasNext(), t)
}

func BenchmarkZip(b *testing.B) {
	Zip(Ints[int](), Ints[int]()).Take(b.N).Consume()
}

func TestEnumerate(t *testing.T) {
	expected := []tuple.T2[int, int]{
		tuple.New2(0, 7),
		tuple.New2(1, 5),
		tuple.New2(2, 3),
		tuple.New2(3, 1),
	}

	test.AssertDeepEq(Enumerate(IntsFromBy(7, -2).Take(4)).Collect(), expected, t)
}

func BenchmarkEnumerate(b *testing.B) {
	Enumerate(Ints[int]()).Take(b.N).Consume()
}

func TestUnzip(t *testing.T) {
	expected := tuple.New2(Ints[int]().Take(10).Collect(), IntsFromBy(10, -1).Take(10).Collect())
	v1, v2 := Unzip(Zip(Elems(expected.V1), Elems(expected.V2)))

	// test.Assert(v1.HasNext(), t)
	// test.Assert(v2.HasNext(), t)

	v1First, _ := v1.Next()
	v2First, _ := v2.Next()
	v2Second, _ := v2.Next()

	test.AssertDeepEq(
		tuple.New2(append([]int{v1First}, v1.Collect()...),
			append([]int{v2First, v2Second}, v2.Collect()...)),
		expected,
		t)
	// test.Assert(!v1.HasNext(), t)
	// test.Assert(!v2.HasNext(), t)
}

func BenchmarkUnzip(b *testing.B) {
	v1, v2 := Unzip(Zip(Ints[int](), Ints[int]()))

	v1.Take(b.N).Consume()
	v2.Take(b.N).Consume()
}
