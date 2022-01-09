package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestTakeWhile(t *testing.T) {
	test.AssertDeepEq(
		InfRange(0, 1).TakeWhile(func(i int) bool { return i < 10 }).Collect(),
		Range(0, 10, 1).Collect(),
		t)
}
