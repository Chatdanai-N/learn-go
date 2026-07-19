package main

import "fmt"

func main() {

	// x := 1
	// for x > 0 {
	// 	if x%2 == 0 {
	// 		fmt.Printf("%v is odd number\n", x)
	// 	}
	// 	x++
	// }

	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			fmt.Printf("%v is odd number\n", i)
		}
	}

}
