package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person
type ByName []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].First < a[j].First
}

func main() {
	p1 := Person{
		First: "James",
		Age:   32,
	}
	p2 := Person{
		First: "Moneypenny",
		Age:   27,
	}
	p3 := Person{
		First: "Q",
		Age:   64,
	}
	p4 := Person{
		First: "M",
		Age:   56,
	}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
