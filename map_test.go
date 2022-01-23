package iter

import (
	"strings"
	"testing"

	"github.com/barweiss/go-tuple"
	"mtoohey.com/iter/test"
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

	// test.Assert(iter.HasNext(), t)
	test.AssertElemsDeepEq(iter.Collect(), expected, t)
	// test.Assert(!iter.HasNext(), t)
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

	test.AssertDeepEq(iter.Collect(), []string{"ITEM1", "ITEM2"}, t)
	// test.Assert(!iter.HasNext(), t)
}

func BenchmarkMapEndoFunc(b *testing.B) {
	Ints[int]().Take(b.N).MapEndo(func(i int) int {
		return i
	}).Consume()
}

func TestMapFunc(t *testing.T) {
	iter := Map(Elems([]string{"item1", "item2"}), func(s string) int { return len(s) })

	test.AssertDeepEq(iter.Collect(), []int{5, 5}, t)
	// test.Assert(!iter.HasNext(), t)
}

func BenchmarkMapFunc(b *testing.B) {
	Map(Ints[int]().Take(b.N), func(i int) int {
		return i
	}).Consume()
}
