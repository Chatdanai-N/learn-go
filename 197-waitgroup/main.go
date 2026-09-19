package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1) // บอก WaitGroup ว่า กำลังจะมี 1 งานใหม่เกิดขึ้น ให้รอด้วย
	go foo()  // สร้าง go routine ใหม่ขึ้นมา
	bar()

	//wg.Wait() //<---- ใส่ตรงนี้เพื่อให้แน่ใจว่า foo() จะทำงานจบด้วย wg.Done() แล้วแน่นอน ไม่ให้เกิด race condition

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() // ใส่ตรงนี้เพื่อที่จะให้เห็นว่า Goroutines ขึ้นมาเป็น 2
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done() // งานเสร็จไปแล้ว 1 งาน
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
