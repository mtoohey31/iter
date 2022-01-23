package iter

import "strings"

// Runes returns an iterator over the runes of the input string.
func Runes(s string) Iter[rune] {
	runes := make([]rune, len(s))
	for i, rune := range s {
		runes[i] = rune
	}
	return Elems(runes)
}

// SplitByRune returns an iterator over the substrings of the input string
// between occurences of the provided rune.
func SplitByRune(s string, r rune) Iter[string] {
	runes := []rune(s)
	index := 0

	return Iter[string](func() (string, bool) {
		newIndex := index
		if len(runes) > index {
			for newIndex < len(runes) {
				if runes[newIndex] == r {
					break
				}
				newIndex++
			}

			res := runes[index:newIndex]
			index = newIndex + 1
			return string(res), true
		} else {
			return "", false
		}
	})
}

type splitByStringInner struct {
	string string
	sep    string
	index  int
}

// SplitByString returns an iterator over the substrings of the input string
// between occurences of the provided separator string.
func SplitByString(s string, sep string) Iter[string] {
	index := 0

	return Iter[string](func() (string, bool) {
		if len(s) > index {
			sepIndex := strings.Index(s[index:], sep)

			if sepIndex == -1 {
				res := s[index:]
				index = len(s)
				return res, true
			} else {
				res := s[index : index+sepIndex]
				index += sepIndex + len(sep)
				return res, true
			}
		} else {
			return "", false
		}
	})
}
