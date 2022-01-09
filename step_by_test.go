package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestStepBy(t *testing.T) {
	test.AssertDeepEq(Range(0, 10, 1).StepBy(3).Collect(), []int{0, 3, 6, 9}, t)
}
