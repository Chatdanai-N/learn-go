package main

import "fmt"

func Logger(f func()) {

	fmt.Println("Starting execution....")
	f()
	fmt.Println("Execution completed.")
}

func main() {

	wrappedFunc := func() {
		fmt.Println("Hello World!")
	}

	Logger(wrappedFunc)

}
