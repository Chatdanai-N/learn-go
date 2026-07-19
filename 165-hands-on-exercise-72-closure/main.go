package main

import (
	"fmt"
	"math"
)

func main() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	y := incrementor2()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())

	fmt.Println("==========================")
	p := powIncrementor(12)
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func incrementor2() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func powIncrementor(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}
