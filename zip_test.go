package iter

import (
	"github.com/barweiss/go-tuple"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestZip(t *testing.T) {
	iter := Zip(Elems([]rune{'a', 'b', 'c', 'd'}), IntsFrom(1))

	expected := []tuple.T2[rune, int]{
		tuple.New2('a', 1),
		tuple.New2('b', 2),
		tuple.New2('c', 3),
		tuple.New2('d', 4),
	}

	assert.Equal(t, iter.Collect(), expected)
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

	assert.Equal(t, Enumerate(IntsFromBy(7, -2).Take(4)).Collect(), expected)
}

func BenchmarkEnumerate(b *testing.B) {
	Enumerate(Ints[int]()).Take(b.N).Consume()
}

func TestUnzip(t *testing.T) {
	expected := tuple.New2(Ints[int]().Take(10).Collect(), IntsFromBy(10, -1).Take(10).Collect())
	v1, v2 := Unzip(Zip(Elems(expected.V1), Elems(expected.V2)))

	v1First, _ := v1()
	v2First, _ := v2()
	v2Second, _ := v2()

	assert.Equal(t, tuple.New2(append([]int{v1First}, v1.Collect()...),
		append([]int{v2First, v2Second}, v2.Collect()...)),
		expected)
}

func BenchmarkUnzip(b *testing.B) {
	v1, v2 := Unzip(Zip(Ints[int](), Ints[int]()))

	v1.Take(b.N).Consume()
	v2.Take(b.N).Consume()
}
