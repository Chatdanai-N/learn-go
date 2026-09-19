package main

import (
	"fmt"
	"mymodule/215-testing-and-benchmarking/03-example-test/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}
