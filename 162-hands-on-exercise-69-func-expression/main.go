package main

import "fmt"

func main() {

	x := func() {
		fmt.Println("Anonymous function")
	}

	y := func(s string) {
		fmt.Println("Say ", s)
	}

	z := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}

	x()
	y("Hello World")
	z()
}
