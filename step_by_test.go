package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStepBy(t *testing.T) {
	iter := Ints[int]().Take(10).StepBy(3)

	assert.Equal(t, []int{0, 3, 6, 9}, iter.Collect())
}

func TestStepByPanic(t *testing.T) {
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
