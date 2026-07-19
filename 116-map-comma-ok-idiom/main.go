package main

import "fmt"

func main() {

	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println("The age of henry was", am["Henry"])
	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	an := make(map[string]int)
	an["Lucus"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	fmt.Println("--- accessing keys that don't exist")
	delete(an, "George")
	fmt.Println(an["George"])
	fmt.Println("-------------------------------------")

	// v, ok := an["George"]
	// if ok {
	// 	fmt.Println("the value prints", v)
	// } else {
	// 	fmt.Println("key didn't exist")
	// }

	if v, ok := an["George"]; ok {
		fmt.Println("the value prints", v)
	} else {
		fmt.Println("key didn't exist")
	}

	// for range over a MAP
	for k, v := range an {
		fmt.Println(k, v)
	}

	for _, v := range an {
		fmt.Println(v)
	}

	for k := range an {
		fmt.Println(k)
	}

	// for range with SLICE

	// xi := []int{42, 43, 44}
	// for k, v := range xi {
	// 	fmt.Println(k, v)
	// }

	// for _, v := range xi {
	// 	fmt.Println(v)
	// }

	// for k := range xi {
	// 	fmt.Println(k)
	// }

}
