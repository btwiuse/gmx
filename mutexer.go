package gmx

// Mutexer is implemented by [Mx] and [RwMx]
type Mutexer[T any] interface {
	// Mutate inner value
	Mut(Mutation[T])
	// Get inner value
	Get() T
}
