package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-flie.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("err happened", err)
		//log.Fatal(err)
		//panic(err)
	}
}
