package main

import "fmt"

func main() {

	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}("Todd")
}

func foo() {
	fmt.Println("Foo ran")
}

// a named function with an indentifier
// func (r receive)  identifier(p parameter(s)) (r return(s)) { code}

// an anonymous function
// func(parameter(s)) (return(s)) { code}
