package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	c := 1
	for i := 0; i < 100; i++ {
		x := RandomNumber(5)

		if x == 3 {
			fmt.Printf("x is %v, total count :%v\n", x, c)
			c++
		}
	}

}

func RandomNumber(n int) int {
	return rand.IntN(n)
}
