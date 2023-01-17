# gmx

Generic Mutex container types for Go, inspired by Rust's
[`Mutex`](https://doc.rust-lang.org/std/sync/struct.Mutex.html) and
[`RwLock`](https://doc.rust-lang.org/std/sync/struct.RwLock.html).

It uses new language features introduced since go1.18:

- generic struct

  - [Mx](https://pkg.go.dev/github.com/btwiuse/gmx#Mx)

  - [RwMx](https://pkg.go.dev/github.com/btwiuse/gmx#RwMx)

- generic interface

  - [Mutexer](https://pkg.go.dev/github.com/btwiuse/gmx#Mutexer)

- generic function

  - [Wrap](https://pkg.go.dev/github.com/btwiuse/gmx#Wrap)

  - [RwWrap](https://pkg.go.dev/github.com/btwiuse/gmx#RwWrap)
