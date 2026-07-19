package main

import (
	"fmt"
	"time"
)

func TimedFunction(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)

	fmt.Println("Elapsed time:", elapsed)
}

func MyFunction() {
	time.Sleep(2 * time.Second)
	fmt.Println("MyFunction is Completed")
}

func main() {

	// Call the warpped function
	TimedFunction(MyFunction)
}
