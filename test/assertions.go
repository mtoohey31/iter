package test

import (
	"reflect"
	"testing"
)

func Assert(v bool, t *testing.T) {
	if !v {
		t.Fatalf("got %v, expected %v", v, true)
	}
}

func AssertEq(actual, expected interface{}, t *testing.T) {
	if actual != expected {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}

func AssertElemsDeepEq[T any](actual, expected []T, t *testing.T) {
	if len(actual) != len(expected) {
		registerAssertElemsDeepEqFail(actual, expected, t)
		return
	}

	uncheckedActual := make(map[int]T)
	for i, a := range actual {
		uncheckedActual[i] = a
	}
	uncheckedExpected := make(map[int]T)
	for i, e := range expected {
		uncheckedExpected[i] = e
	}

	for _, va := range uncheckedActual {
		for ke, ve := range uncheckedExpected {
			if reflect.DeepEqual(va, ve) {
				delete(uncheckedExpected, ke)
			}
		}
	}

	countRemainingExpected := 0
	for range uncheckedExpected {
		countRemainingExpected += 1
	}

	if countRemainingExpected != 0 {
		registerAssertElemsDeepEqFail(actual, expected, t)
		return
	}
}

func registerAssertElemsDeepEqFail[T any](actual T, expected T, t *testing.T) {
	t.Fatalf("got %v, expected the same elements as %v", actual, expected)
}

func AssertDeepEq(actual, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
