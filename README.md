# gmx

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/btwiuse/gmx?tab=doc)
[![Go 1.18+](https://img.shields.io/github/go-mod/go-version/btwiuse/gmx)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/btwiuse/gmx?color=%23000&style=flat-round)](https://github.com/btwiuse/gmx/blob/main/LICENSE)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/btwiuse/gmx)

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
