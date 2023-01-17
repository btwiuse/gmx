package gmx

import "sync"

var _ Mutexer[int] = (*Mx[int])(nil)

// Generic Mutex container type
type Mx[T any] struct {
	*sync.Mutex
	x *T
}

// Wrap a pointer into [Mx]
func Wrap[T any](x *T) *Mx[T] {
	return &Mx[T]{
		Mutex: &sync.Mutex{},
		x:     x,
	}
}

// Do mutable operation on wrapped value guarded by sync.RWMutex
func (m *Mx[T]) Do(op Op[T]) {
	m.Lock()
	defer m.Unlock()
	op(m.x)
}

// Unwrap returns value references by the wrapped pointer
func (m *Mx[T]) Unwrap() T {
	m.Lock()
	defer m.Unlock()
	return *m.x
}
