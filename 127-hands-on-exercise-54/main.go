package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	iceCreamFav []string
}

func main() {

	p1 := person{
		firstName:   "James",
		lastName:    "Bond",
		iceCreamFav: []string{"Chocolate", "Banana", "passion fruit with mango and guava"},
	}

	p2 := person{
		firstName:   "Jenny",
		lastName:    "Moneypenny",
		iceCreamFav: []string{"Mint", "Strawberry", "Chocolate"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.iceCreamFav)
	fmt.Println(p2.iceCreamFav)

	for _, v := range p1.iceCreamFav {
		fmt.Println(p1.firstName, " avorite is ", v)
	}
	for _, v := range p2.iceCreamFav {
		fmt.Println(p2.firstName, " avorite is ", v)
	}

	personMap := make(map[string]person)
	personMap[p1.lastName] = p1
	personMap[p2.lastName] = p2

	for k, v := range personMap {
		fmt.Println("Map key is ", k)
		for _, v2 := range v.iceCreamFav {
			fmt.Println(v.firstName, v.lastName, v2)
		}
	}

}

/*
Take the code from the previous exercise, then store the VALUES of type person in a map with
the KEY of last name. Access each value in the map. Print out the values, ranging over the
slice

*/
