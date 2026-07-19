package main

import "fmt"

func ProcessData(data []int, callback func(int)) {
	for _, item := range data {
		callback(item)
	}
}

func main() {

	data := []int{1, 2, 3, 4, 5}

	callbackFunc := func(num int) {
		fmt.Println("Processing number:", num)
	}

	ProcessData(data, callbackFunc)
}
