package gmx

// Mutation represents mutable operation on a pointer
type Mutation[T any] func(*T)
