package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)
	var incrementor int64
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementor
			runtime.Gosched()
			v++
			incrementor = v
			mu.Unlock()
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("Final Incrementor ", incrementor)
	fmt.Println("Finish !!!")
}
