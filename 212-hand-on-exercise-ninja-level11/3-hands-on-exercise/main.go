package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {

	c := customErr{
		info: "need more coffee",
	}

	foo(c)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
}
