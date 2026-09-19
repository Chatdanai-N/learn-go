package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var gs = 100
	var incrementor int64

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {
			atomic.AddInt64(&incrementor, 1)
			runtime.Gosched()
			atomic.LoadInt64(&incrementor)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Final incrementor :", incrementor)
	fmt.Println("Finish!!!")

}
