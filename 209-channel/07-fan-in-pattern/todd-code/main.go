package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)

}

/*
┌───────────┐
       │  even ───┐│
[send]─┤          ├┼─> [receive] ──> fanin ──> [main print]
       │  odd ────┘│
       └───────────┘

จุดเด่นของโค้ดนี้

Decoupling: ฝั่ง main() ไม่จำเป็นต้องรู้ว่าข้อมูลมาจากกี่แหล่ง (ไม่จำเป็นต้องวนลูปอ่านจาก even และ odd แยกกัน) อ่านแค่ fanin ตัวเดียวพอ

Safe Channel Closing: การใช้ sync.WaitGroup ในฟังก์ชั่น receive
ช่วยป้องกันปัญหา Panic: send on closed channel เพราะทำให้มั่นใจได้ว่า อ่านข้อมูลส่งเข้า fanin จนหมดแล้วจริงๆ ถึงค่อยสั่งปิด fanin

Unordered Output: ลำดับของตัวเลขที่พิมพ์ออกมาอาจไม่เรียง
$0, 1, 2, 3...$ เสมอไป เพราะ even และ odd ทำงานแยกร่างขนานกันใน Goroutine (ขึ้นอยู่กับ Scheduler ของ Go

*/
