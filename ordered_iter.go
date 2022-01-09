package iter

import "github.com/barweiss/go-tuple"

type OrderedIter[T ordered] Iter[T]

func (oi *OrderedIter[T]) Min() (T, error) {
	real := Iter[T](*oi)

	return real.Reduce(func(curr T, next T) T {
		if curr < next {
			return curr
		} else {
			return next
		}
	})
}

func MinByKey[T, U ordered](oi *OrderedIter[T], key func(T) U) (T, error) {
	real := Iter[T](*oi)

	init, err := real.Next()

	if err != nil {
		return real.zeroVal(), IteratorExhaustedError
	}

	return Fold(&real, tuple.New2(init, key(init)), func(curr tuple.T2[T, U], next T) tuple.T2[T, U] {
		keyNext := key(next)
		if curr.V2 < keyNext {
			return curr
		} else {
			return tuple.New2(next, keyNext)
		}
	}).V1, nil
}

func (oi *OrderedIter[T]) Max() (T, error) {
	real := Iter[T](*oi)

	return real.Reduce(func(curr T, next T) T {
		if curr > next {
			return curr
		} else {
			return next
		}
	})
}

func MaxByKey[T, U ordered](oi *OrderedIter[T], key func(T) U) (T, error) {
	real := Iter[T](*oi)

	init, err := real.Next()

	if err != nil {
		return real.zeroVal(), IteratorExhaustedError
	}

	return Fold(&real, tuple.New2(init, key(init)), func(curr tuple.T2[T, U], next T) tuple.T2[T, U] {
		keyNext := key(next)
		if curr.V2 > keyNext {
			return curr
		} else {
			return tuple.New2(next, keyNext)
		}
	}).V1, nil
}

func (oi *OrderedIter[T]) Sum() T {
	real := Iter[T](*oi)

	return real.FoldEndo(real.zeroVal(), func(curr, next T) T { return curr + next })
}
