package gmx

// Mutexer is implemented by [Mx] and [RwMx]
type Mutexer[T any] interface {
	// Mutate inner value
	Do(Mutation[T])
	// Unwrap inner value
	Unwrap() T
}
