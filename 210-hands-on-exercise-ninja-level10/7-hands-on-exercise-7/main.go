package main

import (
	"fmt"
	"runtime"
)

func main() {

	c := make(chan int)
	const goroutine int = 10

	for i := 0; i < goroutine; i++ {

		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			//close(c)
		}()
		fmt.Println("Number of Goroutine:", runtime.NumGoroutine())
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

}

/*
1. เกิด Panic: Send on Closed Channel 💥 (เหตุผลที่สำคัญที่สุด)
ในโค้ดนี้ คุณสั่งรัน Goroutine 10 ตัวพร้อมกัน โดยทุกตัวแชร์ Channel c ร่วมกัน:

for i := 0; i < goroutine; i++ {
    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        // ❌ สมมติว่าสั่ง close(c) ตรงนี้
    }()
}
ปัญหาที่จะเกิดขึ้น: Goroutine ทั้ง 10 ตัวทำงานขนานกัน (Concurrent)
และไม่มีทางทำเสร็จพร้อมกันในวินาทีเดียวกัน

หาก Goroutine ตัวที่ทำเสร็จคนแรก สั่ง close(c) ไปแล้ว
แต่ Goroutine อีก 9 ตัวที่เหลือยังทำไม่เสร็จ และกำลังพยายามส่งเลข c <- i เข้ามา
Go Runtime จะสั่งตัดจบการทำงานด้วยการ Panic ทันที:

panic: send on closed channel


หลักการ Go Design Pattern: "คนส่งหลายคน ห้ามปิด Channel เอง"
Don't close a channel from the receiver side,
 and don't close a channel if the channel has multiple concurrent senders.

ในโจทย์ข้อนี้มี Senders 10 ตัว หากตัวใดตัวหนึ่งชิงปิด Channel จะกระทบกับ Sender ตัวอื่นทันที

หากต้องการปิด Channel นี้จริงๆ จะต้องใช้ sync.WaitGroup เพื่อรอให้ทั้ง 10 Goroutine
ส่งข้อมูลเสร็จทั้งหมดก่อน แล้วค่อยให้ Goroutine กลางตัวอื่นเป็นคนสั่ง close(c) เพียงครั้งเดียว



*/
