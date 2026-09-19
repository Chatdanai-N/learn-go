package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		mutex.Lock()
		counter++
		time.Sleep(time.Duration(rand.IntN(3)) * time.Millisecond)
		fmt.Println(s, i, "Couter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}
