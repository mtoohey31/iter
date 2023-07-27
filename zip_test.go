package iter

import (
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v3/testutils"
)

func FuzzZip(f *testing.F) {
	f.Add([]byte{}, []byte{})
	f.Add([]byte{1, 2}, []byte{3, 4})

	f.Fuzz(func(t *testing.T, a, b []byte) {
		expected := make([]tuple.T2[byte, byte], len(a))

		for i := range a {
			expected[i] = tuple.New2(a[i], b[i])
		}

		assert.Equal(t, expected, Zip(Elems(a), Elems(b)).Collect())
	})
}

func BenchmarkZip(b *testing.B) {
	Zip(Ints[int](), Ints[int]()).Take(uint(b.N)).Consume()
}

func FuzzEnumerate(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		expected := make([]tuple.T2[int, byte], len(b))

		for i, v := range b {
			expected[i] = tuple.New2(i, v)
		}

		assert.Equal(t, expected, Enumerate(Elems(b)).Collect())
	})
}

func BenchmarkEnumerate(b *testing.B) {
	Enumerate(Ints[int]()).Take(uint(b.N)).Consume()
}

func FuzzUnzip(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		i := 0
		iter := func() (tuple.T2[byte, byte], bool) {
			if i < len(b) {
				res := tuple.New2(b[i], b[len(b)-1-i])
				i++
				return res, true
			}

			var z tuple.T2[byte, byte]
			return z, false
		}

		l, r := Unzip(iter)

		bRev := make([]byte, len(b))
		for j, v := range b {
			bRev[len(b)-1-j] = v
		}

		assert.Equal(t, b, l.Collect())
		assert.Equal(t, bRev, r.Collect())

		i = 0
		l, r = Unzip(iter)

		assert.Equal(t, bRev, r.Collect())
		assert.Equal(t, b, l.Collect())
	})
}

func BenchmarkUnzip(b *testing.B) {
	v1, v2 := Unzip(Zip(Ints[int](), Ints[int]()))

	v1.Take(uint(b.N)).Consume()
	v2.Take(uint(b.N)).Consume()
}
