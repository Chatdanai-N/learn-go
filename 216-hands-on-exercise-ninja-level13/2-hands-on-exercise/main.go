package main

import (
	"fmt"
	"mymodule/216-hands-on-exercise-ninja-level13/2-hands-on-exercise/quote"
	"mymodule/216-hands-on-exercise-ninja-level13/2-hands-on-exercise/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
