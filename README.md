# gmx

Generic Mutex container types for Go, inspired by Rust's
[`Mutex`](https://doc.rust-lang.org/std/sync/struct.Mutex.html) and
[`RwLock`](https://doc.rust-lang.org/std/sync/struct.RwLock.html).

It uses new language features introduced since go1.18:

- generic type

  - [Mx](https://pkg.go.dev/github.com/btwiuse/gmx#Mx)

    - [Do](https://pkg.go.dev/github.com/btwiuse/gmx#Mx.Do)

    - [Unwrap](https://pkg.go.dev/github.com/btwiuse/gmx#Mx.Unwrap)

  - [RwMx](https://pkg.go.dev/github.com/btwiuse/gmx#RwMx)

    - [Do](https://pkg.go.dev/github.com/btwiuse/gmx#RwMx.Do)

    - [Unwrap](https://pkg.go.dev/github.com/btwiuse/gmx#RwMx.Unwrap)

  - [Mutation](https://pkg.go.dev/github.com/btwiuse/gmx#Mutation)

- generic interface

  - [Mutexer](https://pkg.go.dev/github.com/btwiuse/gmx#Mutexer)

- generic function

  - [Wrap](https://pkg.go.dev/github.com/btwiuse/gmx#Wrap)

  - [RwWrap](https://pkg.go.dev/github.com/btwiuse/gmx#RwWrap)
