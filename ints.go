package iter

type intsInner[T integer] struct {
	curr T
	by   T
}

func Ints[T integer]() *Iter[T] {
	return WithInner[T](&intsInner[T]{curr: 0, by: 1})
}

func IntsFrom[T integer](start T) *Iter[T] {
	return WithInner[T](&intsInner[T]{curr: start, by: 1})
}

func IntsBy[T integer](by T) *Iter[T] {
	return WithInner[T](&intsInner[T]{by: by})
}

func IntsFromBy[T integer](start T, by T) *Iter[T] {
	return WithInner[T](&intsInner[T]{curr: start, by: by})
}

func (i *intsInner[T]) HasNext() bool { return true }

func (i *intsInner[T]) Next() (T, error) {
	res := i.curr
	i.curr = i.curr + i.by
	return res, nil
}
