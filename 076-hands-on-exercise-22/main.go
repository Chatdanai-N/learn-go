package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	x := RandomNumber(10)
	y := RandomNumber(10)
	fmt.Printf("x and y are %v and %v\t\n", x, y)
	if x < 4 && y < 4 {
		fmt.Printf("%v and %v are both less than 4\n", x, y)
	} else if x > 6 && y > 6 {
		fmt.Printf("%v and %v are both greater than 6\n", x, y)
	} else if x >= 4 && x <= 6 {
		fmt.Printf("%v is greater than or equal to 4 and less than or equal to 6\n", x)
	} else if y != 5 {
		fmt.Printf("%v is not 5\n", y)
	} else {
		fmt.Println("none of the previous were met")
	}

}

func RandomNumber(n int) int {
	return rand.IntN(n)
}
