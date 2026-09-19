package main

import (
	"fmt"
	"mymodule/214-hand-on-exercise-ninja-level12/dog"
)

type canine struct {
	name string
	age  int64
}

func main() {
	fmt.Println("Hello Gopher")

	fido := canine{
		name: "FIdo",
		age:  dog.Years(10),
	}

	fmt.Println(fido)
}
