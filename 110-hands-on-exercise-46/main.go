package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	fmt.Println("--------------------------")

	//inclusive :exclusive
	fmt.Println(x[:3])
	fmt.Println(x[6:])

	fmt.Println("--------------------------")

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
