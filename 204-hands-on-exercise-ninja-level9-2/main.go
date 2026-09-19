package main

import "fmt"

type Person struct {
	first string
	age   int64
}

func (p *Person) speak() {
	fmt.Println("Hello My name is ", p.first, "\t and age", p.age)
}

type Human interface {
	speak()
}

func saySomething(h Human) {
	h.speak()
}

func main() {

	p1 := Person{
		first: "James",
		age:   32,
	}

	p2 := Person{
		first: "Miss Moneypenny",
		age:   27,
	}

	saySomething(&p1)

	p2.speak()

}
