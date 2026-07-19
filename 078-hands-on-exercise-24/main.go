package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	for i := 0; i < 100; i++ {
		x := RandomNumber(10)
		y := RandomNumber(10)
		fmt.Printf("iteration %v \t x and y are %v and %v\t", i, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Printf("both less than 4\n")
		case x > 6 && y > 6:
			fmt.Printf(" both greater than 6\n")
		case x >= 4 && x <= 6:
			fmt.Printf("x is greater than or equal to 4 and less than or equal to 6\n")
		case y != 5:
			fmt.Printf("y is not 5\n")
		default:
			fmt.Println("none of the previous were met")
		}
	}

}

func RandomNumber(n int) int {
	return rand.IntN(n)
}
