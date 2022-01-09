package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestRunes(t *testing.T) {
	test.AssertDeepEq(Runes("asdf").Collect(),
		[]rune{'a', 's', 'd', 'f'}, t)
}

func BenchmarkRunes(b *testing.B) {
	buf := make([]rune, b.N)

	for i := 0; i < b.N; i++ {
		buf[i] = 'a'
	}

	str := string(buf)

	Runes(str).Consume()
}

func TestSplitByRune(t *testing.T) {
	test.AssertDeepEq(SplitByRune("/usr/bin/ls", '/').Collect(),
		[]string{"", "usr", "bin", "ls"}, t)
}

func BenchmarkSplitByRune(b *testing.B) {
	buf := make([]rune, b.N)

	for i := 0; i < b.N; i++ {
		buf[i] = 'a'
	}

	str := string(buf)

	SplitByRune(str, 'a').Consume()
}

func TestSplitByString(t *testing.T) {
	iter := SplitByString("the quick brown fox jumped over the lazy dogs", "the ")

	actual := iter.Collect()
	expected := []string{"", "quick brown fox jumped over ", "lazy dogs"}

	test.AssertDeepEq(actual, expected, t)
}

func BenchmarkSplitByString(b *testing.B) {
	buf := make([]rune, b.N)

	for i := 0; i < b.N; i++ {
		buf[i] = 'a'
	}

	str := string(buf)

	SplitByString(str, "a").Consume()
}
