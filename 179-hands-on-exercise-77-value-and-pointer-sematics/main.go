package main

import "fmt"

type person struct {
	first string
}

func changeFirst(p person, s string) person {
	p.first = s
	return p
}

func changeFirstP(p *person, s string) {
	p.first = s
}

func main() {

	p1 := person{
		first: "James",
	}
	fmt.Println(p1)
	p1 = changeFirst(p1, "James Bond")
	fmt.Println(p1)

	p2 := person{
		first: "Money",
	}
	fmt.Println(p2)
	changeFirstP(&p2, "Money Penny")
	fmt.Println(p2)
}
