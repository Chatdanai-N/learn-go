package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater than ", "10")
	}

	return fmt.Sprint("x is less than ", "10")
}

/*

S1039 - Unnecessary use of fmt.Sprint
Calling fmt.Sprint with a single string argument is unnecessary
and identical to using the string directly.
*/
