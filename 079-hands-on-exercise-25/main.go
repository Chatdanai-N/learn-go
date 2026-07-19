package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	for i := 0; i < 42; i++ {
		x := RandomNumber(5)

		switch x {
		case 0:
			fmt.Printf("iteration %v x is %v\n", i, x)
		case 1:
			fmt.Printf("iteration %v x is %v\n", i, x)
		case 2:
			fmt.Printf("iteration %v x is %v\n", i, x)
		case 3:
			fmt.Printf("iteration %v x is %v\n", i, x)
		case 4:
			fmt.Printf("iteration %v x is %v\n", i, x)
		}
	}

}

func RandomNumber(n int) int {
	return rand.IntN(n)
}
