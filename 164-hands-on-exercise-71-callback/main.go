package main

import (
	"fmt"
)

func main() {

	fmt.Println(printSquare(square, 10))
}

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, n int) string {
	x := f(n)
	return fmt.Sprintf("The Number %v squared is %v", n, x)
}
