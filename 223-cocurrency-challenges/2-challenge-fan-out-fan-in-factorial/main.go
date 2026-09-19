package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	// fan out
	// multiplex multiple channel onto a single channel
	// merge the channels for c0 through c9 onto a single channel
	var y int
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		y++
		fmt.Println(y, "\t", n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100000; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}

		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

// fan in
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	output := func(ch <-chan int) {
		for n := range ch {
			out <- n
		}
		wg.Done()
	}

	for _, c := range cs {
		go output(c)

	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
Challenge #1:
	-- Change the code above to execute 1000 factorial computations concurrently and in parallel.
	-- use the " fan out/ fan in" pattern to accomplish this
Challenge #2:
	-- While running the factorial computations, try to find how much of your resources are being used.const
	-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL

*/
