package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incrementor int64

func main() {

	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementor
			runtime.Gosched()
			v++
			incrementor = v
			wg.Done()
		}()
		//fmt.Println("Incrementor :", incrementor)
		fmt.Println("Number of Goroutine :", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("Final Incrementor :", incrementor)

	fmt.Println("Finish Task !!!")

}
