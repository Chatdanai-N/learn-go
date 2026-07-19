package main

import "fmt"

var x = 40

const y = 41

func main() {

	fmt.Printf("the values of z is %v and the type z is %T\n", x, x)
	fmt.Printf("the values of z is %v and the type z is %T\n", y, y)
	z := 42
	fmt.Printf("the values of z is %v and the type z is %T\n", z, z)
}
