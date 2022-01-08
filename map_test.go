package iter

import (
	"strings"
	"testing"

	"mtoohey.com/iter/test"
)

func TestMapSame(t *testing.T) {
	iter := Elems([]string{"item1", "item2"}).MapSame(func(s string) string { return strings.ToUpper(s) })

	test.AssertDeepEq(iter.Collect(), []string{"ITEM1", "ITEM2"}, t)
	test.Assert(!iter.HasNext(), t)
}

func TestMap(t *testing.T) {
	iter := Map(Elems([]string{"item1", "item2"}), func(s string) int { return len(s) })

	test.AssertDeepEq(iter.Collect(), []int{5, 5}, t)
	test.Assert(!iter.HasNext(), t)
}
