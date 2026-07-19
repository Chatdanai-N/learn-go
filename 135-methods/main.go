package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("i am", p.first)
}

func main() {
	p1 := person{
		first: "James",
	}

	p2 := person{
		first: "Jenny",
	}

	p1.speak()
	p2.speak()
}

// func (r receiver) identifier (p parameter(s)) (return(s)) { code }
