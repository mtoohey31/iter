package iter

import "errors"

type Iter[T any] interface {
	HasNext() bool
	Next() (T, error)
}

var IteratorExhaustedError = errors.New("iterator exhausted")

func Collect[T any](i Iter[T]) []T {
	var res []T
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		res = append(res, next)
	}
	return res
}
