package iter

import (
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v2/testutils"
)

func FuzzKVZip(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		expected := make([]tuple.T2[int, byte], len(b))
		m := make(map[int]byte)
		for i, v := range b {
			m[i] = v
			expected[i] = tuple.New2(i, v)
		}

		assert.ElementsMatch(t, expected, KVZipStrict(m).Collect())
		assert.ElementsMatch(t, expected, KVZipLazy(m).Collect())
	})
}

func BenchmarkKVZipStrict(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	KVZipStrict(m).Consume()
}

func BenchmarkKVZipLazy(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	KVZipLazy(m).Consume()
}

func FuzzMap(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		expected := make([]byte, len(b))
		for i, v := range b {
			expected[i] = v + byte(n)
		}

		assert.Equal(t, expected, Elems(b).Map(func(v byte) byte {
			return v + byte(n)
		}).Collect())
		assert.Equal(t, expected, Map(Elems(b), func(v byte) byte {
			return v + byte(n)
		}).Collect())
	})
}

func BenchmarkIter_Map(b *testing.B) {
	Ints[int]().Take(uint(b.N)).Map(func(i int) int {
		return i
	}).Consume()
}

func BenchmarkMap(b *testing.B) {
	Map(Ints[int]().Take(uint(b.N)), func(i int) int {
		return i
	}).Consume()
}

func FuzzMapWhile(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		var v byte
		var i int
		for i, v = range b {
			if v%2 == 0 {
				break
			}
		}

		iter := Elems(b).MapWhile(func(v byte) (byte, error) {
			if v%2 == 0 {
				return 0, assert.AnError
			}
			return v, nil
		})
		assert.Equal(t, b[:i], iter.Collect())
		_, ok := iter()
		assert.False(t, ok)

		iter = MapWhile(Elems(b), func(v byte) (byte, error) {
			if v%2 == 0 {
				return 0, assert.AnError
			}
			return v, nil
		})
		assert.Equal(t, b[:i], iter.Collect())
		_, ok = iter()
		assert.False(t, ok)
	})
}

func BenchmarkIter_MapWhile(b *testing.B) {
	Ints[int]().Take(uint(b.N)).MapWhile(func(i int) (int, error) {
		return 0, nil
	}).Consume()
}

func BenchmarkMapWhile(b *testing.B) {
	MapWhile(Ints[int]().Take(uint(b.N)), func(i int) (int, error) {
		return 0, nil
	}).Consume()
}

func FuzzFlatMap(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		sum := byte(0)
		for _, v := range b {
			sum += v
		}

		expected := make([]byte, sum)
		i := 0
		for _, v := range b {
			p := i
			for ; i < p+int(v); i++ {
				expected[i] = v
			}
		}

		assert.Equal(t, expected, Elems(b).FlatMap(func(v byte) Iter[byte] {
			return IntsFromBy(v, 0).Take(uint(v))
		}).Collect())
		assert.Equal(t, expected, FlatMap(Elems(b), func(v byte) Iter[byte] {
			return IntsFromBy(v, 0).Take(uint(v))
		}).Collect())
	})
}

func BenchmarkIter_FlatMap_1(b *testing.B) {
	Ints[int]().FlatMap(func(i int) Iter[int] {
		return IntsFrom(i).Take(1)
	}).Take(uint(b.N)).Consume()
}

func BenchmarkIter_FlatMap_100(b *testing.B) {
	Ints[int]().FlatMap(func(i int) Iter[int] {
		return IntsFrom(i).Take(100)
	}).Take(uint(b.N)).Consume()
}

func BenchmarkIter_FlatMap_quarter(b *testing.B) {
	Ints[int]().FlatMap(func(i int) Iter[int] {
		return IntsFrom(i).Take(1 + uint(b.N)/4)
	}).Take(uint(b.N)).Consume()
}

func BenchmarkIter_FlatMap_half(b *testing.B) {
	Ints[int]().FlatMap(func(i int) Iter[int] {
		return IntsFrom(i).Take(1 + uint(b.N)/2)
	}).Take(uint(b.N)).Consume()
}

func BenchmarkIter_FlatMap_full(b *testing.B) {
	Ints[int]().FlatMap(func(i int) Iter[int] {
		return IntsFrom(i).Take(1 + uint(b.N))
	}).Take(uint(b.N)).Consume()
}
