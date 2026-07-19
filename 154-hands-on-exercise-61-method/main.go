package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello my name is %v, and my age is %v\n", p.first, p.age)
}

func main() {

	p1 := person{
		first: "John Snow",
		age:   18,
	}

	p1.speak()
}
