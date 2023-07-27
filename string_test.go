package iter

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v3/internal/testutils"
)

func FuzzRunes(f *testing.F) {
	testutils.AddStrings(f)

	f.Fuzz(func(t *testing.T, s string) {
		assert.Equal(t, []rune(s), Runes(s).Collect())
	})
}

func BenchmarkRunes(b *testing.B) {
	buf := make([]rune, b.N)

	for i := 0; i < b.N; i++ {
		buf[i] = 'a'
	}

	str := string(buf)

	Runes(str).Consume()
}

func FuzzSplitByRune(f *testing.F) {
	f.Add("", ' ')
	f.Add("        ", ' ')
	f.Add("water", '.')
	f.Add("/abs/path", '/')
	f.Add("rel/path", '/')
	f.Add("สวัสดีส", 'ส')
	f.Add("なつ", 'つ')
	f.Add("きたない", 'な')

	f.Fuzz(func(t *testing.T, s string, r rune) {
		if !utf8.ValidRune(r) {
			return
		}

		assert.Equal(t, strings.Split(s, string(r)), SplitByRune(s, r).Collect())
	})
}

func BenchmarkSplitByRune(b *testing.B) {
	buf := make([]rune, b.N)

	for i := 0; i < b.N; i++ {
		buf[i] = 'a'
	}

	str := string(buf)

	SplitByRune(str, 'a').Consume()
}

func FuzzSplitByString(f *testing.F) {
	f.Add("", "")
	f.Add("", " ")
	f.Add("            ", " ")
	f.Add("\xb0", "")
	f.Add(string(utf8.RuneError), "")
	f.Add("water", ".")
	f.Add("[5, 9, 1, 9, 6]", ", ")
	f.Add("the quick brown fox jumped over the lazy dogs", "")
	f.Add("the quick brown fox jumped over the lazy dogs", "the ")

	f.Fuzz(func(t *testing.T, s1, s2 string) {
		assert.Equal(t, strings.Split(s1, s2), SplitByString(s1, s2).Collect())
	})
}

func BenchmarkSplitByString(b *testing.B) {
	buf := make([]rune, b.N)

	for i := 0; i < b.N; i++ {
		buf[i] = 'a'
	}

	str := string(buf)

	SplitByString(str, "a").Consume()
}
