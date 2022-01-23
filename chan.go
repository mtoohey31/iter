package iter

// Receive returns an iterator that reads values from the provided channel, and is
// exhausted when the channel is closed. Note that since this iterator reads
// from a channel, every time the next value is requested the program may end
// up deadlocking if values have not been written: the same rules apply as
// those for reading from a channel in the usual manner.
func Receive[T any](ch *chan T) *Iter[T] {
	tmp := Iter[T](func() (T, bool) {
		next, ok := <-*ch
		return next, ok
	})
	return &tmp
}

// Send consumes the input iterator, sending all yielded values into the
// provided channel. As with receive, this can result in deadlocks if used
// improperly: if nobody is reading from the channel. Also note that this
// method does not close the value after the values have been written, if you
// want that to happen, you should do so yourself.
func (i *Iter[T]) Send(ch *chan T) {
	for {
		next, ok := i.Next()

		if !ok {
			return
		}

		*ch <- next
	}
}
