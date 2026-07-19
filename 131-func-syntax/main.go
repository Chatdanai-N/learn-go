package main

import "fmt"

func main() {

	foo()

	bar("Todd")

	s := aloha("Todd")
	fmt.Println(s)

	s1, age := dogYears("Todd", 4)
	fmt.Println(s1, age)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr.", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, "is this old in dog years"), age
}
