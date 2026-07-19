package main

import "fmt"

func main() {

	defer fmt.Println("Hello World 1")
	defer fmt.Println("Hello World 2")
	fmt.Println("Hello World 3")
}

/*
Hands-on exercise #60 - defer func
● “defer” multiple functions in main
○ show that a deferred func runs after the func containing it exits.
○ determine the order in which the multiple defer funcs run


*/
