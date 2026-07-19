package main

import (
	"fmt"
)

func main() {

	ns1 := []string{"James", "Bond", "Shaken,not stirred"}
	ns2 := []string{"Miss", "Moneypenny", "I'm 008."}

	multins := [][]string{ns1, ns2}
	fmt.Println(multins)

	for _, arrayValue := range multins {
		fmt.Printf("key is %v\n", arrayValue)
		for _, v := range arrayValue {
			fmt.Printf("value is %v\n", v)
		}
	}
}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record.
https://go.dev/play/p/2jxGMyXiZrE */
