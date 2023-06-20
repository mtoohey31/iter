package iter

import (
	"strings"
	"unicode/utf8"
)

// Runes returns an iterator over the runes of the input string.
// Runes(s).Collect() should be equivalent to []rune(s).
func Runes(s string) Iter[rune] {
	b := []byte(s)
	i := 0

	return func() (rune, bool) {
		if len(b) <= i {
			var r rune
			return r, false
		}

		r, size := utf8.DecodeRune(b[i:])
		i += size
		return r, true
	}
}

// SplitByRune returns an iterator over the substrings of the input string
// between occurrences of the provided rune. SplitByRune(s, r).Collect() should
// be equivalent to strings.Split(s, string(r)) for valid runes.
func SplitByRune(s string, r rune) Iter[string] {
	index := 0

	return func() (string, bool) {
		if index < len(s) {
			sepIndex := strings.IndexRune(s[index:], r)

			if sepIndex == -1 {
				res := s[index:]
				index = len(s) + 1
				return res, true
			}

			res := s[index : index+sepIndex]
			index += sepIndex + utf8.RuneLen(r)
			return res, true
		} else if index == len(s) {
			index++
			return "", true
		}

		return "", false
	}
}

// SplitByString returns an iterator over the substrings of the input string
// between occurrences of the provided separator string.
// SplitByString(s, sep).Collect() should be equivalent to
// strings.Split(s, sep).
func SplitByString(s string, sep string) Iter[string] {
	index := 0

	if sep == "" {
		return func() (string, bool) {
			if index < len(s)-1 {
				r, size := utf8.DecodeRuneInString(s[index:])
				prevIndex := index
				index += size

				if r == utf8.RuneError {
					return string(utf8.RuneError), true
				}

				return s[prevIndex:index], true
			} else if index == len(s)-1 {
				index++
				return s[len(s)-1:], true
			}

			return "", false
		}
	}

	return func() (string, bool) {
		if index < len(s) {
			sepIndex := strings.Index(s[index:], sep)

			if sepIndex == -1 {
				res := s[index:]
				index = len(s) + 1
				return res, true
			}

			res := s[index : index+sepIndex]
			index += sepIndex + len(sep)
			return res, true
		} else if index == len(s) {
			index++
			return "", true
		}

		return "", false
	}
}
