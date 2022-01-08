package iter

import "strings"

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
		defer func() { i.index = len(i.string) }()
		return i.string[i.index:], nil
	} else {
		defer func() { i.index = i.index + runeIndex + 1 }()
		return i.string[i.index : i.index+runeIndex], nil
	}
}

type splitByStringInner struct {
	string string
	sep    string
	index  int
}

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
		defer func() { i.index = len(i.string) }()
		return i.string[i.index:], nil
	} else {
		defer func() { i.index = i.index + sepIndex + len(i.sep) }()
		return i.string[i.index : i.index+sepIndex], nil
	}
}
