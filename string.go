package iter

import "strings"

// Runes returns an iterator over the runes of the input string.
func Runes(s string) *Iter[rune] {
	runes := make([]rune, len(s))
	for i, rune := range s {
		runes[i] = rune
	}
	return Elems(runes)
}

type splitByRuneInner struct {
	string string
	rune   rune
	index  int
}

// SplitByRune returns an iterator over the substrings of the input string
// between occurences of the provided rune.
func SplitByRune(s string, r rune) *Iter[string] {
	return WithInner[string](&splitByRuneInner{s, r, 0})
}

func (i *splitByRuneInner) HasNext() bool {
	return i.index < len(i.string)
}

func (i *splitByRuneInner) Next() (string, error) {
	if !i.HasNext() {
		return "", IteratorExhaustedError
	}

	runeIndex := strings.IndexRune(i.string[i.index:], i.rune)

	if runeIndex == -1 {
		res := i.string[i.index:]
		i.index = len(i.string)
		return res, nil
	} else {
		res := i.string[i.index : i.index+runeIndex]
		i.index = i.index + runeIndex + 1
		return res, nil
	}
}

type splitByStringInner struct {
	string string
	sep    string
	index  int
}

// SplitByString returns an iterator over the substrings of the input string
// between occurences of the provided separator string.
func SplitByString(s string, sep string) *Iter[string] {
	return WithInner[string](&splitByStringInner{s, sep, 0})
}

func (i *splitByStringInner) HasNext() bool {
	return i.index < len(i.string)
}

func (i *splitByStringInner) Next() (string, error) {
	if !i.HasNext() {
		return "", IteratorExhaustedError
	}

	sepIndex := strings.Index(i.string[i.index:], i.sep)

	if sepIndex == -1 {
		res := i.string[i.index:]
		i.index = len(i.string)
		return res, nil
	} else {
		res := i.string[i.index : i.index+sepIndex]
		i.index = i.index + sepIndex + len(i.sep)
		return res, nil
	}
}
