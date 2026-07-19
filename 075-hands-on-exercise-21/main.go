package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is from the init function")
}

func main() {

	n := RandomNumber(400)

	switch {
	case n <= 100:
		fmt.Printf("%v between 0 and 100", n)
	case n > 101 && n <= 200:
		fmt.Printf("%v between 101 and 200 ", n)
	case n >= 200 && n <= 250:
		fmt.Printf("%v between 201 and 250", n)
	default:
		fmt.Printf("%v This was more than 250", n)
	}

}

func RandomNumber(n int) int {
	return rand.Intn(n)
}
