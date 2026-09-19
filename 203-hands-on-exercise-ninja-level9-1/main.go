package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Number of CPUs\t", runtime.NumCPU())
	fmt.Println("Number of Goroutiner\t", runtime.NumGoroutine())
	wg.Add(2)

	go func() {
		fmt.Println("Hello print something one !")
		fmt.Println("Number of Goroutiner\t", runtime.NumGoroutine())
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello print something two !")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Number of Goroutiner\t", runtime.NumGoroutine())
	fmt.Println("about to exits")
}

/*
Hands-on exercise #1
● in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists

*/
