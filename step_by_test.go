package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestStepBy(t *testing.T) {
	iter := Ints[int]().Take(10).StepBy(3)

	// test.Assert(iter.HasNext(), t)
	test.AssertDeepEq(iter.Collect(), []int{0, 3, 6, 9}, t)
	// test.Assert(!iter.HasNext(), t)
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
