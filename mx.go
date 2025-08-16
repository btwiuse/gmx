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

// Mutate wrapped value guarded by sync.RWMutex
func (m *Mx[T]) Mut(op Mutation[T]) {
	m.Lock()
	defer m.Unlock()
	op(m.x)
}

// Get value referenced by the wrapped pointer
func (m *Mx[T]) Get() T {
	m.Lock()
	defer m.Unlock()
	return *m.x
}
