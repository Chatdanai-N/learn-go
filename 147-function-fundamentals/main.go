package main

import "fmt"

func main() {
	foo()
}

// func (r receiver) identifer(p parameter(s)) (r return(s)) { code } ( argument(s))

func foo() {
	fmt.Println("Foo ran")
}
