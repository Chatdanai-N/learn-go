package main

import (
	"fmt"
)

type person struct {
	first     string
	friends   map[string]int
	favDrinks []string
}

func main() {

	p1 := person{
		first: "James",
		friends: map[string]int{
			"Robert": 1,
		},
		favDrinks: []string{"Water", "Beer", "Pepsi"},
	}

	p2 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Penny",
		friends: map[string]int{
			"Robert": 1,
			"Jane":   2,
		},
		favDrinks: []string{"Water", "Milk"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for k, v := range p2.friends {
		fmt.Println(p2.first, " - friends - ", k, v)
	}

	for _, v := range p2.favDrinks {
		fmt.Println(p2.first, " - drinks - ", v)
	}
}

/*
Create and use an anonymous struct with these fields:
● first string
● friends map[string]int
● favDrinks []string

*/
