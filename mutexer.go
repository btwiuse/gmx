package gmx

// Mutexer is implemented by [Mx] and [RwMx]
type Mutexer[T any] interface {
	// Mutate inner value
	Do(Op[T])
	// Unwrap inner value
	Unwrap() T
}
