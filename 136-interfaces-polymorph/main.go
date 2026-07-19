package main

import "fmt"

type person struct {
	first string
}

type secertAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secertAgent) speak() {
	fmt.Println("I am secert agent", sa.first)
}

func main() {

	sa1 := secertAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	p2 := person{
		first: "Jenny",
	}

	sa1.speak()
	p2.speak()
}
