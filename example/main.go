package main

import (
	"sync"

	"github.com/btwiuse/gmx"
)

func main() {
	MxExample()
	RwMxExample()
}

func MxExample() {
	// initialize x
	var x int = 0
	println(x)

	// initialize mx
	var m *gmx.Mx[*int] = gmx.Wrap(&x)
	println(*m.Unwrap())

	// initialize wait group
	var wg *sync.WaitGroup = &sync.WaitGroup{}

	// initialize mutation op
	var op gmx.Op[*int] = func(x *int) {
		*x += 1
		wg.Done()
	}

	// with gmx.Do: 100000
	{
		wg.Add(100000)
		for i := 0; i < 100000; i++ {
			go m.Do(op)
		}
		wg.Wait()

		println(*m.Unwrap())
	}

	// without gmx.Do: 193287
	{
		wg.Add(100000)
		for i := 0; i < 100000; i++ {
			go op(&x)
		}
		wg.Wait()

		println(*m.Unwrap())
	}
}

func RwMxExample() {
	// initialize x
	var x int = 0
	println(x)

	// initialize rwmx
	var m *gmx.RwMx[*int] = gmx.RwWrap(&x)
	println(*m.Unwrap())

	// initialize wait group
	var wg *sync.WaitGroup = &sync.WaitGroup{}

	// initialize mutation op
	var op gmx.Op[*int] = func(x *int) {
		*x += 1
		wg.Done()
	}

	// with rwmx.Do: 100000
	{
		wg.Add(100000)
		for i := 0; i < 100000; i++ {
			go m.Do(op)
		}
		wg.Wait()

		println(*m.Unwrap())
	}

	// without rwmx.Do: 193287
	{
		wg.Add(100000)
		for i := 0; i < 100000; i++ {
			go op(&x)
		}
		wg.Wait()

		println(*m.Unwrap())
	}
}
