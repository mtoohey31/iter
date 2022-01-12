package iter

type intsInner[T integer] struct {
	curr T
	by   T
}

// Ints returns an iterator that produces integer values of the specified
// generic type, starting from 0 and increasing by 1.
func Ints[T integer]() *Iter[T] {
	return Wrap[T](&intsInner[T]{curr: 0, by: 1})
}

// IntsFrom returns an iterator that produces integer values of the specified
// generic type, starting from the provided value and increasing by 1.
func IntsFrom[T integer](start T) *Iter[T] {
	return Wrap[T](&intsInner[T]{curr: start, by: 1})
}

// IntsBy returns an iterator that produces integer values of the specified
// generic type, starting from 0 and increasing by the provided value.
func IntsBy[T integer](by T) *Iter[T] {
	return Wrap[T](&intsInner[T]{by: by})
}

// IntsFromBy returns an iterator that produces integer values of the specified
// generic type, starting from, and increasing by the provided values.
func IntsFromBy[T integer](start T, by T) *Iter[T] {
	return Wrap[T](&intsInner[T]{curr: start, by: by})
}

func (i *intsInner[T]) HasNext() bool { return true }

func (i *intsInner[T]) Next() (T, error) {
	res := i.curr
	i.curr = i.curr + i.by
	return res, nil
}
