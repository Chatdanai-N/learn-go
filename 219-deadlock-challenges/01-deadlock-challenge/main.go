package main

import "fmt"

func main() {
	c := make(chan int)
	// fix deadlock
	go func() {
		c <- 1
	}()

	fmt.Println(<-c)
}
