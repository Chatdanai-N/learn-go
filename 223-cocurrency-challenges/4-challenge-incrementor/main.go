package main

import (
	"fmt"
	"strconv"
	"sync"
)

var count int64
var wg sync.WaitGroup

func main() {

	//old version
	// wg.Add(2)
	// go incrementor("1")
	// go incrementor("2")
	// wg.Wait()

	c := incrementor(2)
	var count int

	for n := range c {
		count++
		fmt.Println(n)
	}

	fmt.Println("Final Counter:", count)
}

// old version
// func incrementor(s string) {
// 	for i := 0; i < 20; i++ {
// 		atomic.AddInt64(&count, 1)
// 		fmt.Println("Process:"+s+" printing:", i)
// 	}
// 	wg.Done()
// }

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprint("Process:"+strconv.Itoa(i)+" printing:", k)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}

/*
Challenge #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
