package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {

	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wished", "Never say nerver"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))

}

func toJSON(p1 person) ([]byte, error) {
	bs, err := json.Marshal(p1)
	if err != nil {
		return []byte{}, fmt.Errorf("there was an error in toJSON: %v", err)
		//return []byte{}, errors.New(fmt.Sprintf("There was an error in toJSON", err))
	}
	return bs, nil
}
