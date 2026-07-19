package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {

	p1 := person{
		firstName:   "James",
		lastName:    "Bond",
		favIceCream: []string{"Chocolate", "Banana", "passion fruit with mango and guava"},
	}

	p2 := person{
		firstName:   "Jenny",
		lastName:    "Moneypenny",
		favIceCream: []string{"Mint", "Strawberry", "Chocolate"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.favIceCream)
	fmt.Println(p2.favIceCream)

	for _, v := range p1.favIceCream {
		fmt.Println(p1.firstName, " avorite is ", v)
	}
	for _, v := range p2.favIceCream {
		fmt.Println(p2.firstName, " avorite is ", v)
	}
}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.

*/
