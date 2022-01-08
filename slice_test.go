package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestArrIter(t *testing.T) {
	expected := []string{"item1", "item2"}
	test.AssertDeepEq(Elems(expected).Collect(), expected, t)
}
