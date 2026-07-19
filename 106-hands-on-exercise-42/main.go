package main

import "fmt"

func main() {

	arrayNumber := [5]int{}

	for i := 0; i < 5; i++ {
		arrayNumber[i] = i
	}

	for i, v := range arrayNumber {
		fmt.Printf("index %v : value %v and type %T\n", i, v, v)
	}

}
