package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestRunes(t *testing.T) {
	test.AssertDeepEq(Runes("asdf").Collect(),
		[]rune{'a', 's', 'd', 'f'}, t)
}

func TestSplitByRune(t *testing.T) {
	test.AssertDeepEq(SplitByRune("/usr/bin/ls", '/').Collect(),
		[]string{"", "usr", "bin", "ls"}, t)
}

func TestSplitByString(t *testing.T) {
	iter := SplitByString("the quick brown fox jumped over the lazy dogs", "the ")

	actual := iter.Collect()
	expected := []string{"", "quick brown fox jumped over ", "lazy dogs"}

	test.AssertDeepEq(actual, expected, t)
}
