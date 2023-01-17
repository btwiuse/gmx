package gmx

// Op represents mutable operation on a pointer
type Op[T any] func(x *T)
