package main

type Pausable interface {
	pause()
	resume()
}

type Frozen[T Pausable] struct {
	T *Pausable // use this to save while paused
}
