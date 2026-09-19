package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go putNumbers(c1)

	go getAllNumber(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("about to exits")

}

func putNumbers(c1 chan<- int) {

	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)
}

func getAllNumber(c1, c2 chan int) {

	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)

		go func(v2 int) {
			c2 <- v2
			wg.Done()
		}(v)

	}

	wg.Wait()
	close(c2)
}
