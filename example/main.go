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
		batchIncrMutexer(m, 1e6).x, // with gmx.Mut: 100000
		batchIncrPointer(s, 1e6).x, // without gmx.Mut: 93287
	)
}

func TestRwMx() {
	var (
		s *state           = &state{x: 0}  // initialize state
		m *gmx.RwMx[state] = gmx.RwWrap(s) // initialize rwmx
	)

	println(
		"gmx.RwMx",
		batchIncrMutexer(m, 1e6).x, // with rwmx.Mut: 100000
		batchIncrPointer(s, 1e6).x, // without rwmx.Mut: 99939
	)
}

func incrOp() gmx.Mutation[state] {
	return func(s *state) {
		s.x += 1
	}
}

// use Mutexer.Mut to perform state mutation is safe
func batchIncrMutexer(m gmx.Mutexer[state], n int) state {
	var (
		wg = &sync.WaitGroup{} // initialize wait group
		op = incrOp()          // initialize incr op
	)

	for i := 0; i < n; i++ {
		wg.Go(func() { m.Mut(op) })
	}

	wg.Wait()

	return m.Get()
}

// use pointer to perform state mutation is unsafe
func batchIncrPointer(s *state, n int) state {
	var (
		wg = &sync.WaitGroup{} // initialize wait group
		op = incrOp()          // initialize incr op
	)

	for i := 0; i < n; i++ {
		wg.Go(func() { op(s) })
	}

	wg.Wait()

	return *s
}
