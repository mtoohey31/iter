package iter

import (
	"errors"
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
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

		assert.ElementsMatch(t, expected, KVZip(m).Collect())
		assert.ElementsMatch(t, expected, KVZipChannelled(m).Collect())
	})
}

func BenchmarkKVZip(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	KVZip(m).Consume()
}

func BenchmarkKVZipChannelled(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	KVZipChannelled(m).Consume()
}

func FuzzMap(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		expected := make([]byte, len(b))
		for i, v := range b {
			expected[i] = v + byte(n)
		}

		assert.Equal(t, expected, Elems(b).MapEndo(func(v byte) byte {
			return v + byte(n)
		}).Collect())
		assert.Equal(t, expected, Map(Elems(b), func(v byte) byte {
			return v + byte(n)
		}).Collect())
	})
}

func BenchmarkMapEndoFunc(b *testing.B) {
	Ints[int]().Take(b.N).MapEndo(func(i int) int {
		return i
	}).Consume()
}

func BenchmarkMapFunc(b *testing.B) {
	Map(Ints[int]().Take(b.N), func(i int) int {
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

		iter := Elems(b).MapWhileEndo(func(v byte) (byte, error) {
			if v%2 == 0 {
				return 0, errors.New("")
			}
			return v, nil
		})
		assert.Equal(t, b[:i], iter.Collect())
		_, ok := iter()
		assert.False(t, ok)

		iter = MapWhile(Elems(b), func(v byte) (byte, error) {
			if v%2 == 0 {
				return 0, errors.New("")
			}
			return v, nil
		})
		assert.Equal(t, b[:i], iter.Collect())
		_, ok = iter()
		assert.False(t, ok)
	})
}

func BenchmarkMapWhileEndo(b *testing.B) {
	Ints[int]().Take(b.N).MapWhileEndo(func(i int) (int, error) {
		return 0, nil
	}).Consume()
}

func BenchmarkMapWhile(b *testing.B) {
	MapWhile(Ints[int]().Take(b.N), func(i int) (int, error) {
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

		assert.Equal(t, expected, Elems(b).FlatMapEndo(func(v byte) Iter[byte] {
			return IntsFromBy(v, 0).Take(int(v))
		}).Collect())
		assert.Equal(t, expected, FlatMap(Elems(b), func(v byte) Iter[byte] {
			return IntsFromBy(v, 0).Take(int(v))
		}).Collect())
	})
}

func BenchmarkFlatMapEndo1(b *testing.B) {
	Ints[int]().FlatMapEndo(func(i int) Iter[int] {
		return IntsFrom(i).Take(1)
	}).Take(b.N).Consume()
}

func BenchmarkFlatMapEndo100(b *testing.B) {
	Ints[int]().FlatMapEndo(func(i int) Iter[int] {
		return IntsFrom(i).Take(100)
	}).Take(b.N).Consume()
}

func BenchmarkFlatMapEndoQuarter(b *testing.B) {
	Ints[int]().FlatMapEndo(func(i int) Iter[int] {
		return IntsFrom(i).Take(1 + b.N/4)
	}).Take(b.N).Consume()
}

func BenchmarkFlatMapEndoHalf(b *testing.B) {
	Ints[int]().FlatMapEndo(func(i int) Iter[int] {
		return IntsFrom(i).Take(1 + b.N/2)
	}).Take(b.N).Consume()
}

func BenchmarkFlatMapEndoFull(b *testing.B) {
	Ints[int]().FlatMapEndo(func(i int) Iter[int] {
		return IntsFrom(i).Take(1 + b.N)
	}).Take(b.N).Consume()
}
