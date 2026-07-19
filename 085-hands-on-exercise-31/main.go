package main

import "fmt"

func main() {

	m := map[string]int{
		"Jamnes":     42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("Key: %v, value %v\n", k, v)
	}
}
