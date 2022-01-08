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

func AssertDeepEq(actual, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v, expected %v", actual, expected)
	}
}
