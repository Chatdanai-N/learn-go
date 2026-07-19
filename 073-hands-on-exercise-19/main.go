package main

import (
	"fmt"
	"math/rand"
)

func main() {

	n := RandomNumber(400)

	if n <= 100 {
		fmt.Printf("%v between 0 and 100", n)
	} else if n >= 101 && n <= 200 {
		fmt.Printf("%v between 101 and 200 ", n)
	} else if n >= 201 && n <= 250 {
		fmt.Printf("%v between 201 and 250", n)
	} else {
		fmt.Printf("%v This was more than 250", n)
	}

}

func RandomNumber(n int) int {
	return rand.Intn(n)
}
