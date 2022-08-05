package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzStepBy(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		m := int(n + 1)
		expected := []byte{}
		for i := 0; i < len(b); i += m {
			expected = append(expected, b[i])
		}

		assert.Equal(t, expected, Elems(b).StepBy(m).Collect())
	})
}

func TestStepBy_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("StepBy should've panicked")
		}
	}()

	Elems([]bool{}).StepBy(0)
}

func BenchmarkStepBy1(b *testing.B) {
	Ints[int]().StepBy(1).Take(b.N).Consume()
}

func BenchmarkStepBy100(b *testing.B) {
	Ints[int]().StepBy(100).Take(b.N).Consume()
}
