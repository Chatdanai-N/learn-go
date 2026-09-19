package main

import (
	"fmt"
	"mymodule/216-hands-on-exercise-ninja-level13/1-hands-on-exercise/dog"
)

type canine struct {
	name string
	age  int
}

func main() {

	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}

	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
