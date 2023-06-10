package iter

// Cycle creates an iterator that first consumes the provided input iterator,
// then repeatedly returns the prior values. The second return value is true
// if the input iterator was non-empty (meaning the first return value is an
// infinite iterator) or false otherwise (meaning the first return value is
// nil).
func (i Iter[T]) Cycle() (Iter[T], bool) {
	next, ok := i()
	if !ok {
		return nil, false
	}

	cachedFirst := &next
	memory := []T{*cachedFirst}
	index := -1

	var self Iter[T]
	self = func() (T, bool) {
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
			}

			index = 0
			return self()
		}

		res := memory[index]
		index = (index + 1) % len(memory)
		return res, true

	}
	return self, true
}
