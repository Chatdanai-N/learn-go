package main

import "fmt"

func main() {

	x, y, z := 747, 911, 90210
	fmt.Printf("binary:%b \t\t decimal:%d \t hexadecimal:%#X\n", x, x, x)
	fmt.Printf("binary:%b \t\t decimal:%d \t hexadecimal:%#X\n", y, y, y)
	fmt.Printf("binary:%b \t decimal:%d \t hexadecimal:%#X\n", z, z, z)
}
