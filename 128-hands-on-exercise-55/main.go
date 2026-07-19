package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  int
	color  string
}

func main() {

	bydAtto := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "BYD",
		model: "ATTO 3",
		doors: 4,
		color: "Black",
	}

	hondaCity := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "HONDA",
		model: "CITY NEW",
		doors: 4,
		color: "White",
	}

	fmt.Println(bydAtto)
	fmt.Println(hondaCity)

	fmt.Println(bydAtto.engine.electric, bydAtto.make, bydAtto.model)
	fmt.Println(hondaCity.engine.electric, hondaCity.make, hondaCity.model)
}

/*
● Create a type engine struct, and include this field
○ electric bool
● Create a type vehicle struct, and include these fields
■ engine
■ make
■ model
■ doors
■ color
● Create two VALUES of TYPE vehicle
○ use a composite literal
● Print out each of these values.
● Print out a single field from each of these values.

*/
