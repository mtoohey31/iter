package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunes(t *testing.T) {
	assert.Equal(t, []rune{'a', 's', 'd', 'f'}, Runes("asdf").Collect())
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
	assert.Equal(t, []string{"", "usr", "bin", "ls"},
		SplitByRune("/usr/bin/ls", '/').Collect())
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

	assert.Equal(t, expected, actual)
}

func BenchmarkSplitByString(b *testing.B) {
	buf := make([]rune, b.N)

	for i := 0; i < b.N; i++ {
		buf[i] = 'a'
	}

	str := string(buf)

	SplitByString(str, "a").Consume()
}
