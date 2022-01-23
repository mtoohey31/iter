package iter

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapWhileEndo(t *testing.T) {
	initialIter := Elems([]string{"good", "bad", "good", "good"})
	mappedWhileIter := initialIter.MapWhileEndo(func(s string) (string, error) {
		if s == "bad" {
			return "", errors.New("")
		} else {
			return strings.ToUpper(s), nil
		}
	})

	assert.Equal(t, mappedWhileIter.Collect(), []string{"GOOD"})
	assert.Equal(t, initialIter.Collect(), []string{"good", "good"})

	Ints[int]().Take(5).MapWhileEndo(func(i int) (int, error) {
		return i, nil
	}).Consume()
}

func BenchmarkMapWhileEndo(b *testing.B) {
	Ints[int]().Take(b.N).MapWhileEndo(func(i int) (int, error) {
		return 0, nil
	}).Consume()
}

func TestMapWhile(t *testing.T) {
	initialIter := Elems([]string{"long string", "longer string", "short", "long string again"})
	mappedWhileIter := MapWhile(initialIter, func(s string) (int, error) {
		l := len(s)
		if l < 10 {
			return 0, errors.New("")
		} else {
			return l, nil
		}
	})

	assert.Equal(t, mappedWhileIter.Collect(), []int{11, 13})

	_, ok := mappedWhileIter()

	assert.False(t, ok)
	assert.Equal(t, initialIter.Collect(), []string{"long string again"})
}

func BenchmarkMapWhile(b *testing.B) {
	MapWhile(Ints[int]().Take(b.N), func(i int) (int, error) {
		return 0, nil
	}).Consume()
}
