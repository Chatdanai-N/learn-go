package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous function")
	}()

	func(s string) {
		fmt.Println("Say ", s)
	}("Hello World")

	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
}
