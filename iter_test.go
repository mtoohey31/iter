package iter

import (
	"errors"
	"sync"
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v2/testutils"
)

func FuzzIter_Collect(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		assert.Equal(t, b, Elems(b).Collect())
	})
}

func BenchmarkIter_Collect(b *testing.B) {
	Ints[int]().Take(uint(b.N)).Collect()
}

func FuzzIter_CollectInto(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		actual := make([]byte, n)
		var expectedN = len(b)
		if int(n) < expectedN {
			expectedN = int(n)
		}
		assert.Equal(t, expectedN, Elems(b).CollectInto(actual))
		var expected []byte
		if int(n) < len(b) {
			expected = b[:n]
		} else {
			expected = append(b, make([]byte, int(n)-len(b))...)
		}
		assert.Equal(t, expected, actual)
	})
}

func BenchmarkIter_CollectInto(b *testing.B) {
	slice := make([]int, b.N)
	Ints[int]().CollectInto(slice)
}

func FuzzIter_All(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{1})
	f.Add([]byte{1, 2, 3, 4})
	f.Add([]byte{1, 3, 5, 7})
	f.Add([]byte{0, 2, 4, 6, 8})

	f.Fuzz(func(t *testing.T, b []byte) {
		expected := true
		for _, v := range b {
			if v%2 != 0 {
				expected = false
			}
		}

		assert.Equal(t, expected, Elems(b).All(func(v byte) bool { return v%2 == 0 }))
	})
}

func BenchmarkIter_All(b *testing.B) {
	Ints[int]().Take(uint(b.N)).All(func(i int) bool {
		return i >= 0
	})
}

func FuzzIter_Any(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{1})
	f.Add([]byte{1, 2, 3, 4})
	f.Add([]byte{1, 3, 5, 7})
	f.Add([]byte{0, 2, 4, 6, 8})

	f.Fuzz(func(t *testing.T, b []byte) {
		expected := false
		for _, v := range b {
			if v%2 == 0 {
				expected = true
			}
		}

		assert.Equal(t, expected, Elems(b).Any(func(v byte) bool { return v%2 == 0 }))
	})
}

func BenchmarkIter_Any(b *testing.B) {
	Ints[int]().Take(uint(b.N)).Any(func(i int) bool {
		return i < 0
	})
}

func FuzzIter_Count(f *testing.F) {
	testutils.AddUints(f)

	f.Fuzz(func(t *testing.T, n uint) {
		assert.Equal(t, int(n), Ints[int]().Take(n).Count())
	})
}

func BenchmarkIter_Count(b *testing.B) {
	Ints[int]().Take(uint(b.N)).Count()
}

func FuzzIter_Find(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		actual, ok := Zip(Elems(b), Ints[uint]()).Find(
			func(t tuple.T2[byte, uint]) bool {
				return t.V2 == n
			})
		if assert.Equal(t, uint(len(b)) > n, ok) && ok {
			assert.Equal(t, n, actual.V2)
		}
	})
}

func FuzzFindMap(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		actual, ok := Zip(Elems(b), Ints[uint]()).FindMap(
			func(t tuple.T2[byte, uint]) (tuple.T2[byte, uint], error) {
				if t.V2 == n {
					return tuple.New2(t.V1+1, t.V2+1), nil
				}

				return tuple.New2(byte(0), uint(0)), errors.New("")
			})
		if assert.Equal(t, uint(len(b)) > n, ok) && ok {
			assert.Equal(t, n+1, actual.V2)
		}

		actual, ok = FindMap(Zip(Elems(b), Ints[uint]()),
			func(t tuple.T2[byte, uint]) (tuple.T2[byte, uint], error) {
				if t.V2 == n {
					return tuple.New2(t.V1+1, t.V2+1), nil
				}

				return tuple.New2(byte(0), uint(0)), errors.New("")
			})
		if assert.Equal(t, uint(len(b)) > n, ok) && ok {
			assert.Equal(t, n+1, actual.V2)
		}
	})
}

func FuzzFold(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		expected := byte(0)
		for _, v := range b {
			expected += v
		}

		assert.Equal(t, expected, Elems(b).Fold(0, func(sum, v byte) byte {
			return sum + v
		}))
		assert.Equal(t, expected, Fold(Elems(b), 0, func(sum, v byte) byte {
			return sum + v
		}))
	})
}

func BenchmarkIter_Fold(b *testing.B) {
	Ints[int]().Take(uint(b.N)).Fold(0, func(p, n int) int {
		return p + n
	})
}

func BenchmarkFold(b *testing.B) {
	Fold(Ints[int]().Take(uint(b.N)), 0, func(p, n int) int {
		return p + n
	})
}

func FuzzIter_ForEach(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		var expected byte
		for _, v := range b {
			expected += v
		}

		var actual byte
		Elems(b).ForEach(func(v byte) { actual += v })
		assert.Equal(t, expected, actual)

		var m sync.Mutex
		actual = 0
		Elems(b).ForEachParallel(func(v byte) {
			m.Lock()
			defer m.Unlock()
			actual += v
		})
		assert.Equal(t, expected, actual)
	})
}

func BenchmarkIter_ForEach(b *testing.B) {
	Ints[int]().Take(uint(b.N)).ForEach(func(i int) {})
}

func BenchmarkIter_ForEachParallel(b *testing.B) {
	Ints[int]().Take(uint(b.N)).ForEachParallel(func(i int) {})
}

func FuzzIter_Last(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		actual, ok := Elems(b).Last()
		if len(b) == 0 {
			assert.False(t, ok)
			assert.Zero(t, actual)
		} else if assert.True(t, ok) {
			assert.Equal(t, b[len(b)-1], actual)
		}
	})
}

func BenchmarkIter_Last(b *testing.B) {
	Ints[int]().Take(uint(b.N)).Last()
}

func FuzzIter_Nth(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		actual, ok := Elems(b).Nth(int(n))
		if len(b) <= int(n) {
			assert.False(t, ok)
			assert.Zero(t, actual)
		} else if assert.True(t, ok) {
			assert.Equal(t, b[n], actual)
		}
	})
}

func BenchmarkIter_Nth(b *testing.B) {
	Ints[int]().Nth(b.N)
}

func FuzzTryFold(f *testing.F) {
	err := errors.New("")

	f.Add(5, true)
	f.Add(5, false)
	f.Add(18, true)
	f.Add(18, false)
	f.Add(318, true)
	f.Add(318, false)

	f.Fuzz(func(t *testing.T, n int, b bool) {
		expected := 0
		for i := 0; i < n; i++ {
			expected += i
		}

		actual, actualErr := Ints[int]().Take(uint(n)).TryFold(0, func(sum, v int) (int, error) {
			if b {
				return 0, err
			}

			return sum + v, nil
		})

		if b {
			assert.Same(t, err, actualErr)
		} else {
			if assert.NoError(t, actualErr) {
				assert.Equal(t, expected, actual)
			}
		}

		actual, actualErr = TryFold(Ints[int]().Take(uint(n)), 0, func(sum, v int) (int, error) {
			if b {
				return 0, err
			}

			return sum + v, nil
		})

		if b {
			assert.True(t, err == actualErr)
		} else {
			if assert.NoError(t, actualErr) {
				assert.Equal(t, expected, actual)
			}
		}
	})
}

func BenchmarkIter_TryFold(b *testing.B) {
	Ints[int]().Take(uint(b.N)).TryFold(0, func(curr, next int) (int, error) {
		return 0, nil
	})
}

func BenchmarkTryFold(b *testing.B) {
	TryFold(Ints[int]().Take(uint(b.N)), 0, func(curr, next int) (int, error) {
		return 0, nil
	})
}

func FuzzIter_TryForEach(f *testing.F) {
	err := errors.New("")

	f.Add(5, true)
	f.Add(5, false)
	f.Add(18, true)
	f.Add(18, false)
	f.Add(318, true)
	f.Add(318, false)

	f.Fuzz(func(t *testing.T, n int, b bool) {
		expected := 0
		for i := 0; i < n; i++ {
			expected += i
		}

		actual := 0
		actualErr := Ints[int]().Take(uint(n)).TryForEach(func(v int) error {
			if b {
				return err
			}

			actual += v
			return nil
		})

		if b {
			assert.Same(t, err, actualErr)
		} else {
			if assert.NoError(t, actualErr) {
				assert.Equal(t, expected, actual)
			}
		}
	})
}

func BenchmarkIter_TryForEach(b *testing.B) {
	Ints[int]().Take(uint(b.N)).TryForEach(func(i int) error { return nil })
}

func FuzzIter_Reduce(f *testing.F) {
	testutils.AddUints(f)

	f.Fuzz(func(t *testing.T, n uint) {
		expected := 1
		for i := 0; i < int(n); i++ {
			expected *= i
		}

		actual, ok := Ints[int]().Take(n).Reduce(func(prod, v int) int {
			return prod * v
		})

		if n == 0 {
			assert.False(t, ok)
		} else {
			assert.Equal(t, expected, actual)
		}
	})
}

func BenchmarkIter_Reduce(b *testing.B) {
	Ints[int]().Take(uint(b.N)).Reduce(func(p, n int) int {
		return 0
	})
}

func FuzzIter_Pos(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		iter := Elems(b)
		for _, v1 := range b {
			if !assert.Zero(t, iter.Pos(func(v2 byte) bool {
				return v1 == v2
			})) {
				break
			}
		}
		assert.Equal(t, -1, iter.Pos(func(byte) bool { return true }))
	})
}

func BenchmarkIter_Pos(b *testing.B) {
	Ints[int]().Pos(func(i int) bool { return i == b.N })
}

func FuzzIter_Rev(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		expected := make([]byte, len(b))
		for i, v := range b {
			expected[len(b)-1-i] = v
		}

		assert.Equal(t, expected, Elems(b).Rev().Collect())
	})
}

func BenchmarkIter_Rev(b *testing.B) {
	Ints[int]().Take(uint(b.N)).Rev()
}
