package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestStepBy(t *testing.T) {
	test.AssertDeepEq(Ints[int]().Take(10).StepBy(3).Collect(), []int{0, 3, 6, 9}, t)
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
