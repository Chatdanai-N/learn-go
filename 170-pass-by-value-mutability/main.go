package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(xi []int) {
	xi[0] = 99
}

func mapDelta(mp map[string]int, k string) {
	mp[k] = 33
}

func main() {

	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])
}
