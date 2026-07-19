package main

import "fmt"

func main() {

	// nameMap := map[string][]string{
	// 	"bond_james":       []string{"shaken, not stirred", "martinis", "fast cars"},
	// 	"moneypenny_jenny": []string{"intelligence", "literature", "computer science"},
	// 	"no_dr":            []string{"cats", "ice cream", "sunsets"},
	// }

	nameMap := make(map[string][]string)
	nameMap["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	nameMap["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	nameMap["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	for k, v := range nameMap {
		fmt.Println("Key is : ", k)
		// fmt.Println("length of value ", len(v))
		// fmt.Println("values is ", v)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
