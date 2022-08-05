package iter

import "strings"

// Runes returns an iterator over the runes of the input string.
func Runes(s string) Iter[rune] {
	return Elems([]rune(s))
}

// SplitByRune returns an iterator over the substrings of the input string
// between occurences of the provided rune.
func SplitByRune(s string, r rune) Iter[string] {
	// TODO: this isn't lazy!
	runes := []rune(s)
	index := 0

	return func() (string, bool) {
		newIndex := index
		if index < len(runes) {
			for newIndex < len(runes) {
				if runes[newIndex] == r {
					break
				}
				newIndex++
			}

			res := runes[index:newIndex]
			index = newIndex + 1
			return string(res), true
		} else if index == len(runes) {
			index += 1
			return "", true
		} else {
			return "", false
		}
	}
}

// SplitByString returns an iterator over the substrings of the input string
// between occurences of the provided separator string.
func SplitByString(s string, sep string) Iter[string] {
	index := 0

	return func() (string, bool) {
		if index < len(s) {
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
	}
}
