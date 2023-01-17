package gmx

import "sync"

// Generic Read-Write Mutex container type
type RwMx[T any] struct {
	*sync.RWMutex
	x T
}

// Wrap a value, usually a pointer, into [RwMx]
func RwWrap[T any](x T) *RwMx[T] {
	return &RwMx[T]{
		RWMutex: &sync.RWMutex{},
		x:       x,
	}
}

// Do mutable operation on wrapped value guarded by sync.RWMutex
func (m *RwMx[T]) Do(op Op[T]) {
	m.Lock()
	defer m.Unlock()
	op(m.x)
}

// Unwrap wrapped value, usually a pointer
func (m *RwMx[T]) Unwrap() T {
	m.RLock()
	defer m.RUnlock()
	return m.x
}
