package main

import (
	"sync"

	"github.com/btwiuse/gmx"
)

func main() {
	TestMx()
	TestRwMx()
}

type state struct {
	x int
}

func TestMx() {
	var (
		s *state         = &state{x: 0} // initialize state
		m *gmx.Mx[state] = gmx.Wrap(s)  // initialize mx
	)

	println(
		"gmx.Mx  ",
		batchIncrMutexer(m, 1e6).x, // with gmx.Do: 100000
		batchIncrPointer(s, 1e6).x, // without gmx.Do: 93287
	)
}

func TestRwMx() {
	var (
		s *state           = &state{x: 0}  // initialize state
		m *gmx.RwMx[state] = gmx.RwWrap(s) // initialize rwmx
	)

	println(
		"gmx.RwMx",
		batchIncrMutexer(m, 1e6).x, // with rwmx.Do: 100000
		batchIncrPointer(s, 1e6).x, // without rwmx.Do: 99939
	)
}

func incrOp(wg *sync.WaitGroup) gmx.Mutation[state] {
	return func(s *state) {
		s.x += 1
		wg.Done()
	}
}

// use Mutexer.Do to perform state mutation is safe
func batchIncrMutexer(m gmx.Mutexer[state], n int) state {
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

// use pointer to perform state mutation is unsafe
func batchIncrPointer(s *state, n int) state {
	var (
		wg = &sync.WaitGroup{} // initialize wait group
		op = incrOp(wg)        // initialize incr op
	)

	wg.Add(n)

	for i := 0; i < n; i++ {
		go op(s)
	}

	wg.Wait()

	return *s
}
