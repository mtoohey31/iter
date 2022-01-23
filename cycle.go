package iter

// Cycle returns an iterator that first consumes the provided input iterator,
// then repeatedly returns the previous values. It also returns a boolean that
// indicates whether the creation of this iterator was successful: it will fail
// if the provided iterator is already empty.
func (i Iter[T]) Cycle() (Iter[T], bool) {
	next, ok := i()
	if !ok {
		return nil, false
	}

	cachedFirst := &next
	memory := []T{*cachedFirst}
	index := -1

	var self Iter[T]
	self = Iter[T](func() (T, bool) {
		if index == -1 {
			if cachedFirst != nil {
				res := *cachedFirst
				cachedFirst = nil
				return res, true
			}

			next, ok := i()
			if ok {
				memory = append(memory, next)
				return next, true
			} else {
				index = 0
				return self()
			}
		} else {
			res := memory[index]
			index = (index + 1) % len(memory)
			return res, true
		}
	})
	return self, true
}
