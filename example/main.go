package main

import (
	"sync"

	"github.com/btwiuse/gmx"
)

func main() {
	TestMx()
	TestRwMx()
}

func TestMx() {
	var (
		x int                         // initialize x
		m *gmx.Mx[int] = gmx.Wrap(&x) // initialize mx
	)

	println(
		"gmx.Mx  ",
		incrMutexer(m, 1e6),  // with gmx.Do: 100000
		incrPointer(&x, 1e6), // without gmx.Do: 93287
	)
}

func TestRwMx() {
	var (
		x int                             // initialize x
		m *gmx.RwMx[int] = gmx.RwWrap(&x) // initialize rwmx
	)

	println(
		"gmx.RwMx",
		incrMutexer(m, 1e6),  // with rwmx.Do: 100000
		incrPointer(&x, 1e6), // without rwmx.Do: 99939
	)
}

func incrOp(wg *sync.WaitGroup) gmx.Op[int] {
	return func(x *int) {
		*x += 1
		wg.Done()
	}
}

func incrMutexer(m gmx.Mutexer[int], n int) int {
	var (
		wg = &sync.WaitGroup{} // initialize wait group
		op = incrOp(wg)        // initialize incr op
	)

	wg.Add(n)

	for i := 0; i < n; i++ {
		go m.Do(op)
	}

	wg.Wait()

	return m.Unwrap()
}

func incrPointer(x *int, n int) int {
	var (
		wg = &sync.WaitGroup{} // initialize wait group
		op = incrOp(wg)        // initialize incr op
	)

	wg.Add(n)

	for i := 0; i < n; i++ {
		go op(x)
	}

	wg.Wait()

	return *x
}
