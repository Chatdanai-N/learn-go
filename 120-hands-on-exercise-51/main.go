package main

import "fmt"

func main() {

	nameMap := make(map[string][]string)
	nameMap["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	nameMap["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	nameMap["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	nameMap["fleming_ian"] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range nameMap {
		fmt.Println("Key is : ", k)
		// fmt.Println("length of value ", len(v))
		// fmt.Println("values is ", v)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

	// delete record in map
	fmt.Println("============record deleted==================")
	delete(nameMap, "fleming_ian")
	fmt.Println("============record deleted==================")
	for k, v := range nameMap {
		fmt.Println("Key is : ", k)
		// fmt.Println("length of value ", len(v))
		// fmt.Println("values is ", v)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
