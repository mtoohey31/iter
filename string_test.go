package iter

import (
	"reflect"
	"testing"
)

func TestRunes(t *testing.T) {
	iter := Runes("asdf")

	actual := iter.Collect()
	expected := []rune{'a', 's', 'd', 'f'}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func TestSplitByRune(t *testing.T) {
	iter := SplitByRune("/usr/bin/ls", '/')

	actual := iter.Collect()
	expected := []string{"", "usr", "bin", "ls"}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func TestSplitByString(t *testing.T) {
	iter := SplitByString("the quick brown fox jumped over the lazy dogs", "the ")

	actual := iter.Collect()
	expected := []string{"", "quick brown fox jumped over ", "lazy dogs"}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
