package gmx

import "sync"

var _ Mutexer[int] = (*RwMx[int])(nil)

// Generic Read-Write Mutex container type
type RwMx[T any] struct {
	*sync.RWMutex
	x *T
}

// Wrap a pointer into [RwMx]
func RwWrap[T any](x *T) *RwMx[T] {
	return &RwMx[T]{
		RWMutex: &sync.RWMutex{},
		x:       x,
	}
}

// Mutate wrapped value guarded by sync.RWMutex
func (m *RwMx[T]) Mut(op Mutation[T]) {
	m.Lock()
	defer m.Unlock()
	op(m.x)
}

// Get value referenced by the wrapped pointer
func (m *RwMx[T]) Get() T {
	m.RLock()
	defer m.RUnlock()
	return *m.x
}
