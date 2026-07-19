package main

import "fmt"

func main() {

	numberSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range numberSlice {
		fmt.Printf("index %v, value is %v and type %T\n", i, v, v)
	}

	fmt.Printf("%T \t %#v\n", numberSlice, numberSlice)

}
