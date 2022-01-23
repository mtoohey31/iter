package iter

import (
	"errors"
	"strings"
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/assert"
)

func TestMapData(t *testing.T) {
	expected := []tuple.T2[string, int]{
		tuple.New2("1", 1),
		tuple.New2("2", 2),
		tuple.New2("3", 3),
		tuple.New2("4", 4),
	}

	m := make(map[string]int)
	for _, v := range expected {
		m[v.V1] = v.V2
	}

	iter := KVZip(m)

	assert.ElementsMatch(t, iter.Collect(), expected)
}

func BenchmarkMapData(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	KVZip(m).Consume()
}

func TestMapEndoFunc(t *testing.T) {
	iter := Elems([]string{"item1", "item2"}).MapEndo(func(s string) string { return strings.ToUpper(s) })

	assert.Equal(t, iter.Collect(), []string{"ITEM1", "ITEM2"})
}

func BenchmarkMapEndoFunc(b *testing.B) {
	Ints[int]().Take(b.N).MapEndo(func(i int) int {
		return i
	}).Consume()
}

func TestMapFunc(t *testing.T) {
	iter := Map(Elems([]string{"item1", "item2"}), func(s string) int { return len(s) })

	assert.Equal(t, iter.Collect(), []int{5, 5})
}

func BenchmarkMapFunc(b *testing.B) {
	Map(Ints[int]().Take(b.N), func(i int) int {
		return i
	}).Consume()
}

func TestMapWhileEndo(t *testing.T) {
	initialIter := Elems([]string{"good", "bad", "good", "good"})
	mappedWhileIter := initialIter.MapWhileEndo(func(s string) (string, error) {
		if s == "bad" {
			return "", errors.New("")
		} else {
			return strings.ToUpper(s), nil
		}
	})

	assert.Equal(t, mappedWhileIter.Collect(), []string{"GOOD"})
	assert.Equal(t, initialIter.Collect(), []string{"good", "good"})

	Ints[int]().Take(5).MapWhileEndo(func(i int) (int, error) {
		return i, nil
	}).Consume()
}

func BenchmarkMapWhileEndo(b *testing.B) {
	Ints[int]().Take(b.N).MapWhileEndo(func(i int) (int, error) {
		return 0, nil
	}).Consume()
}

func TestMapWhile(t *testing.T) {
	initialIter := Elems([]string{"long string", "longer string", "short", "long string again"})
	mappedWhileIter := MapWhile(initialIter, func(s string) (int, error) {
		l := len(s)
		if l < 10 {
			return 0, errors.New("")
		} else {
			return l, nil
		}
	})

	assert.Equal(t, mappedWhileIter.Collect(), []int{11, 13})

	_, ok := mappedWhileIter()

	assert.False(t, ok)
	assert.Equal(t, initialIter.Collect(), []string{"long string again"})
}

func BenchmarkMapWhile(b *testing.B) {
	MapWhile(Ints[int]().Take(b.N), func(i int) (int, error) {
		return 0, nil
	}).Consume()
}

func TestFlatMapEndo(t *testing.T) {
	initial := []int{1, 2, 3}
	iter := Elems(initial).FlatMapEndo(func(i int) Iter[int] {
		return IntsFrom(i).Take(2)
	})

	actual := iter.Collect()
	expected := []int{1, 2, 2, 3, 3, 4}

	assert.Equal(t, actual, expected)
}

func TestFlatMap(t *testing.T) {
	initial := []string{"alpha", "beta", "gamma"}
	iter := FlatMap(Elems(initial), func(s string) Iter[rune] {
		return Runes(s)
	})

	actualStart := iter.Take(5).Collect()
	expected := strings.Join(initial, "")

	assert.Equal(t, string(append(actualStart, iter.Collect()...)), expected)
}

//
// func BenchmarkFlatMapEndo1(b *testing.B) {
// 	InfRange(0, 1).FlatMapEndo(func(i int) Iter[int] {
// 		return Range(i, i+1, 1)
// 	}).Take(b.N).Consume()
// }
//
// func BenchmarkFlatMapEndo100(b *testing.B) {
// 	InfRange(0, 1).FlatMapEndo(func(i int) Iter[int] {
// 		return Range(i, i+100, 1)
// 	}).Take(b.N).Consume()
// }
//
// func BenchmarkFlatMapEndoQuarter(b *testing.B) {
// 	InfRange(0, 1).FlatMapEndo(func(i int) Iter[int] {
// 		return Range(i, b.N/4, 1)
// 	}).Take(b.N).Consume()
// }
//
// func BenchmarkFlatMapEndoHalf(b *testing.B) {
// 	InfRange(0, 1).FlatMapEndo(func(i int) Iter[int] {
// 		return Range(i, b.N/2, 1)
// 	}).Take(b.N).Consume()
// }
//
// func BenchmarkFlatMapEndoFull(b *testing.B) {
// 	InfRange(0, 1).FlatMapEndo(func(i int) Iter[int] {
// 		return Range(i, b.N, 1)
// 	}).Take(b.N).Consume()
// }
