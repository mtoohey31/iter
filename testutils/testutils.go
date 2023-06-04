package testutils

import (
	"testing"
	"unicode/utf8"
)

func AddByteSlicePairs(f *testing.F) {
	f.Add([]byte{}, []byte{})
	f.Add([]byte{}, []byte{1, 2})
	f.Add([]byte{1, 2}, []byte{})
	f.Add([]byte{1, 2}, []byte{3, 4})
}

func AddByteSliceUintPairs(f *testing.F) {
	f.Add([]byte{}, uint(0))
	f.Add([]byte{}, uint(13))
	f.Add([]byte{1, 2, 3, 4, 5, 6}, uint(0))
	f.Add([]byte{1, 2, 3, 4, 5, 6}, uint(3))
	f.Add([]byte{1, 2, 3, 4, 5, 6}, uint(48))
}

func AddByteSlices(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{1, 2, 3, 4})
	f.Add([]byte{1, 1, 2, 3, 5, 8, 13, 21, 34, 55})
}

func AddStrings(f *testing.F) {
	f.Add("")
	f.Add("the")
	f.Add("quick brown fox")
	f.Add("よる")
	f.Add("こんにちは")
	f.Add("สวัสดีส")
	f.Add("\xb0")
	f.Add(string(utf8.RuneError))
	f.Add("/")
	f.Add("/absolute/path")
	f.Add("relative/path")
}

func AddUints(f *testing.F) {
	f.Add(uint(0))
	f.Add(uint(27))
	f.Add(uint(168))
	f.Add(uint(41981))
}
