package iter

type chainInner[T any] struct {
	firstInner  *Iter[T]
	secondInner *Iter[T]
}

func (i *Iter[T]) Chain(o *Iter[T]) *Iter[T] {
	return WithInner[T](&chainInner[T]{firstInner: i, secondInner: o})
}

func (i *chainInner[T]) HasNext() bool {
	return i.firstInner.HasNext() || i.secondInner.HasNext()
}

func (i *chainInner[T]) Next() (T, error) {
	if i.firstInner.HasNext() {
		return i.firstInner.Next()
	} else {
		return i.secondInner.Next()
	}
}
