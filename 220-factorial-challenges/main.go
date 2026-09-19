package main

import "fmt"

func main() {
	f := factorial(4)

	fmt.Println("Total:", <-f)
	for n := range f {
		fmt.Println(n)
	}
}

// func factorial(n int) int {
// 	total := 1
// 	for i := n; i > 0; i-- {
// 		total *= i
// 	}
// 	return total
// }

func factorial(n int) chan int {
	out := make(chan int)

	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()

	return out
}

/*
CHALLENGE #1
-- Use gorouties and channel to calculate factorial

CHALLENGE #2
-- Why might you want to use goroutine and channels to calculate factorial ?
-- read a few fo the other answers at the discussion area to see the reasons of other
*/
