package gmx

// Op is an operation on some value x, usually a pointer
type Op[T any] func(x T)
