package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {

	s := `[{"first":"James","last":"Bond","age":32},{"first":"Miss","last":"Moneypenny","age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var person []person
	err := json.Unmarshal(bs, &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", person)

	for i, v := range person {
		fmt.Println("---- person number:", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
