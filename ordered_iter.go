package iter

import "github.com/barweiss/go-tuple"

func Min[T ordered](oi *Iter[T]) (T, error) {
	return oi.Reduce(func(curr T, next T) T {
		if curr < next {
			return curr
		} else {
			return next
		}
	})
}

func MinByKey[T, U ordered](oi *Iter[T], key func(T) U) (T, error) {
	init, err := oi.Next()

	if err != nil {
		return oi.zeroVal(), IteratorExhaustedError
	}

	return Fold(oi, tuple.New2(init, key(init)), func(curr tuple.T2[T, U], next T) tuple.T2[T, U] {
		keyNext := key(next)
		if curr.V2 < keyNext {
			return curr
		} else {
			return tuple.New2(next, keyNext)
		}
	}).V1, nil
}

func Max[T ordered](oi *Iter[T]) (T, error) {
	return oi.Reduce(func(curr T, next T) T {
		if curr > next {
			return curr
		} else {
			return next
		}
	})
}

func MaxByKey[T, U ordered](oi *Iter[T], key func(T) U) (T, error) {
	init, err := oi.Next()

	if err != nil {
		return oi.zeroVal(), IteratorExhaustedError
	}

	return Fold(oi, tuple.New2(init, key(init)), func(curr tuple.T2[T, U], next T) tuple.T2[T, U] {
		keyNext := key(next)
		if curr.V2 > keyNext {
			return curr
		} else {
			return tuple.New2(next, keyNext)
		}
	}).V1, nil
}

func Sum[T ordered](oi *Iter[T]) T {
	return oi.FoldEndo(oi.zeroVal(), func(curr, next T) T { return curr + next })
}
