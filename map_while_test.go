package iter

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestMapWhileSame(t *testing.T) {
	initialIter := Elems([]string{"good", "bad", "good", "good"})
	mappedWhileIter := initialIter.MapWhileSame(func(s string) (string, error) {
		if s == "bad" {
			return "", errors.New("")
		} else {
			return strings.ToUpper(s), nil
		}
	})

	actualMappedWhile := mappedWhileIter.Collect()
	expectedMappedWhile := []string{"GOOD"}

	if !reflect.DeepEqual(actualMappedWhile, expectedMappedWhile) {
		t.Fatalf("got %v, expected %v", actualMappedWhile, expectedMappedWhile)
	}

	actualInitial := initialIter.Collect()
	expectedInitial := []string{"good", "good"}

	if !reflect.DeepEqual(actualInitial, expectedInitial) {
		t.Fatalf("got %v, expected %v", actualInitial, expectedInitial)
	}
}

func TestMapwhile(t *testing.T) {
	initialIter := Elems([]string{"long string", "longer string", "short", "long string again"})
	mappedWhileIter := MapWhile(initialIter, func(s string) (int, error) {
		l := len(s)
		if l < 10 {
			return 0, errors.New("")
		} else {
			return l, nil
		}
	})

	actualMappedWhile := mappedWhileIter.Collect()
	expectedMappedWhile := []int{11, 13}

	if !reflect.DeepEqual(actualMappedWhile, expectedMappedWhile) {
		t.Fatalf("got %v, expected %v", actualMappedWhile, expectedMappedWhile)
	}

	actualInitial := initialIter.Collect()
	expectedInitial := []string{"long string again"}

	if !reflect.DeepEqual(actualInitial, expectedInitial) {
		t.Fatalf("got %v, expected %v", actualInitial, expectedInitial)
	}
}
