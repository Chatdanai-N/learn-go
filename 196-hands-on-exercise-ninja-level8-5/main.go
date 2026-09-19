package main

import (
	"fmt"
	"sort"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByLast []User
type ByAge []User

func (l ByLast) Len() int {
	return len(l)
}

func (l ByLast) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ByLast) Less(i, j int) bool {
	return l[i].Last < l[j].Last
}

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {

	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}
	fmt.Println(users)
	// your code goes here
	// sort by last
	sort.Sort(ByLast(users))
	fmt.Println("Sort By Lastname============")
	for _, user := range users {
		fmt.Println("FirstName : ", user.First, ", Lastname : ", user.Last, ", Age :", user.Age)
		for _, saying := range user.Sayings {
			fmt.Println("\t\t", saying)
		}

	}

	// sort by age
	sort.Sort(ByAge(users))
	fmt.Println("Sort By Age===========")
	for _, user := range users {
		fmt.Println("FirstName : ", user.First, ", Lastname : ", user.Last, ", Age :", user.Age)
		for _, saying := range user.Sayings {
			fmt.Println("\t\t", saying)
		}

	}

}
