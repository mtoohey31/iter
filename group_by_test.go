package iter

import (
	"math/rand"
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mtoohey.com/iter/v3/internal/testutils"
)

func groupBySlice[K comparable, V any](s []V, f func(V) K) []tuple.T2[K, []V] {
	res := []tuple.T2[K, []V]{}

	if len(s) == 0 {
		return res
	}

	currentKey := f(s[0])
	groupSoFar := []V{s[0]}
	for _, v := range s[1:] {
		nextKey := f(v)
		if nextKey != currentKey {
			res = append(res, tuple.T2[K, []V]{V1: currentKey, V2: groupSoFar})
			currentKey = nextKey
			groupSoFar = nil
		}

		groupSoFar = append(groupSoFar, v)
	}

	return append(res, tuple.T2[K, []V]{V1: currentKey, V2: groupSoFar})
}

// smush sticks the two slices together, preserving the original orderings of
// elements as they are arranged in the input slices, but interleaving elements
// of the two slices amongst each other randomly.
func smush[T any](a []T, b []T) []T {
	res := make([]T, len(a)+len(b))

	for i := 0; ; i++ {
		if len(a) == 0 {
			copy(res[i:], b)
			break
		}
		if len(b) == 0 {
			copy(res[i:], a)
			break
		}

		if rand.Int()%2 == 0 {
			res[i] = a[0]
			a = a[1:]
		} else {
			res[i] = b[0]
			b = b[1:]
		}
	}

	return res
}

func FuzzGroupBy(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		groupFunc := func(b byte) bool { return b%2 == 0 }

		expectedPartitions := groupBySlice(b, groupFunc)

		t.Run("sub-iterators evaluated as we go", func(t *testing.T) {
			oi := GroupBy(Elems(b), groupFunc)

			for _, expectedPartition := range expectedPartitions {
				p, ok := oi()

				require.True(t, ok)
				assert.Equal(t, expectedPartition.V1, p.V1)
				assert.Equal(t, expectedPartition.V2, p.V2.Collect())
			}

			p, ok := oi()

			if assert.False(t, ok) {
				assert.Zero(t, p)
			}
		})

		t.Run("sub-iterators evaluated later", func(t *testing.T) {
			actualSubIterators := []Iter[byte]{}

			oi := GroupBy(Elems(b), groupFunc)

			for _, expectedPartition := range expectedPartitions {
				p, ok := oi()

				require.True(t, ok)
				assert.Equal(t, expectedPartition.V1, p.V1)
				actualSubIterators = append(actualSubIterators, p.V2)
			}

			p, ok := oi()

			if assert.False(t, ok) {
				assert.Zero(t, p)
			}

			require.Len(t, actualSubIterators, len(expectedPartitions))

			for i, expectedPartition := range expectedPartitions {
				assert.Equal(t, expectedPartition.V2, actualSubIterators[i].Collect())
			}
		})

		t.Run("sub-iterators evaluated ぐちゃぐちゃ", func(t *testing.T) {
			operations := []func(){}

			oi := GroupBy(Elems(b), groupFunc)

			for i, expectedPartition := range expectedPartitions {
				i := i
				expectedPartition := expectedPartition
				operations = append(operations, func() {
					t.Logf("evaluating oi index %d, expecting: %#v", i, expectedPartition)

					p, ok := oi()

					require.True(t, ok)
					assert.Equal(t, expectedPartition.V1, p.V1)

					subIteratorOperations := []func(){}
					for j, expectedNext := range expectedPartition.V2 {
						j := j
						expectedNext := expectedNext
						subIteratorOperations = append(subIteratorOperations, func() {
							t.Logf("evaluating si index %d,%d, expecting: %#v", i, j, expectedNext)

							next, ok := p.V2()

							if assert.True(t, ok) {
								assert.Equal(t, expectedNext, next)
							}
						})
					}
					subIteratorOperations = append(subIteratorOperations, func() {
						t.Logf("evaluating si index %d,%d, expecting nothing", i, len(expectedPartition.V2))

						next, ok := p.V2()

						if assert.False(t, ok) {
							assert.Zero(t, next)
						}
					})
					operations = smush(operations, subIteratorOperations)
				})
			}
			operations = append(operations, func() {
				t.Logf("evaluating oi index %d, expecting nothing", len(expectedPartitions))

				next, ok := oi()

				if assert.False(t, ok) {
					assert.Zero(t, next)
				}
			})

			for len(operations) > 0 {
				operation := operations[0]
				operations = operations[1:]
				operation()
			}
		})
	})
}

func BenchmarkGroupBy_asWeGo_1(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) bool {
		return v%2 == 0
	})

	for next, ok := oi(); ok; next, ok = oi() {
		next.V2.Consume()
	}
}

func BenchmarkGroupBy_asWeGo_10(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) bool {
		return (v/10)%2 == 0
	})

	for next, ok := oi(); ok; next, ok = oi() {
		next.V2.Consume()
	}
}

func BenchmarkGroupBy_asWeGo_100(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) bool {
		return (v/100)%2 == 0
	})

	for next, ok := oi(); ok; next, ok = oi() {
		next.V2.Consume()
	}
}

func BenchmarkGroupBy_asWeGo_quarter(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) bool {
		return (v*4/b.N)%2 == 0
	})

	for next, ok := oi(); ok; next, ok = oi() {
		next.V2.Consume()
	}
}

func BenchmarkGroupBy_asWeGo_half(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) bool {
		return (v*2/b.N)%2 == 0
	})

	for next, ok := oi(); ok; next, ok = oi() {
		next.V2.Consume()
	}
}

func BenchmarkGroupBy_asWeGo_full(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) struct{} {
		return struct{}{}
	})

	for next, ok := oi(); ok; next, ok = oi() {
		next.V2.Consume()
	}
}

func BenchmarkGroupBy_later_1(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) bool {
		return v%2 == 0
	})

	sis := []Iter[int]{}
	for next, ok := oi(); ok; next, ok = oi() {
		sis = append(sis, next.V2)
	}

	for _, si := range sis {
		si.Consume()
	}
}

func BenchmarkGroupBy_later_10(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) bool {
		return (v/10)%2 == 0
	})

	sis := []Iter[int]{}
	for next, ok := oi(); ok; next, ok = oi() {
		sis = append(sis, next.V2)
	}

	for _, si := range sis {
		si.Consume()
	}
}

func BenchmarkGroupBy_later_100(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) bool {
		return (v/100)%2 == 0
	})

	sis := []Iter[int]{}
	for next, ok := oi(); ok; next, ok = oi() {
		sis = append(sis, next.V2)
	}

	for _, si := range sis {
		si.Consume()
	}
}

func BenchmarkGroupBy_later_quarter(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) bool {
		return (v*4/b.N)%2 == 0
	})

	sis := []Iter[int]{}
	for next, ok := oi(); ok; next, ok = oi() {
		sis = append(sis, next.V2)
	}

	for _, si := range sis {
		si.Consume()
	}
}

func BenchmarkGroupBy_later_half(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) bool {
		return (v*2/b.N)%2 == 0
	})

	sis := []Iter[int]{}
	for next, ok := oi(); ok; next, ok = oi() {
		sis = append(sis, next.V2)
	}

	for _, si := range sis {
		si.Consume()
	}
}

func BenchmarkGroupBy_later_full(b *testing.B) {
	oi := GroupBy(Ints[int]().Take(uint(b.N)), func(v int) struct{} {
		return struct{}{}
	})

	sis := []Iter[int]{}
	for next, ok := oi(); ok; next, ok = oi() {
		sis = append(sis, next.V2)
	}

	for _, si := range sis {
		si.Consume()
	}
}
