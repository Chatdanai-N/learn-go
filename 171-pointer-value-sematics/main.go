package main

import "fmt"

func addOne(v int) int {
	return v + 1
}

func addOneP(v *int) {
	*v += 1
}

func main() {

	// value sematics
	a := 1
	fmt.Println("Value semantics ===")
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)

	//pointer semantics
	b := 1
	fmt.Println("Pointer semantics ===")
	fmt.Println(b) // 1
	addOneP(&b)
	fmt.Println(b) // 2
}
