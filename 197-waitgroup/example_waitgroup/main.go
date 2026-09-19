package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// บอก WaitGroup ว่าจะทำทั้งหมด 5 งาน
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		// สั่งรัน Goroutine 5 ตัวแยกกัน
		go worker(i)
	}

	// รอให้ตัวนับลดลงจนเหลือ 0 (ครบทั้ง 5 งาน)
	wg.Wait()
	fmt.Println("ทุกงานเสร็จสิ้นแล้ว! จบโปรแกรม")
}

func worker(id int) {
	defer wg.Done() // เมื่อฟังก์ชันนี้ทำงานเสร็จ จะสั่งลดตัวนับลง 1 งานอัตโนมัติ

	fmt.Printf("Worker %d: กำลังเริ่มทำงาน...\n", id)
	time.Sleep(time.Second) // จำลองการทำงานหนัก 1 วินาที
	fmt.Printf("Worker %d: เสร็จงานแล้ว!\n", id)
}
