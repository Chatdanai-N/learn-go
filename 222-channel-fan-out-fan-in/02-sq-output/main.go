package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen(2, 3)

	// FAN OUT
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n) //4 then 9 or 9 then4
	}

}

func gen(nums ...int) chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
			   ┌──> [ sq (c1) ] ──┐
[ gen (2, 3) ] ┤                  ├─> [ merge ] ──> [ main (print) ]
               └──> [ sq (c2) ] ──┘
               │                  │
           (FAN OUT)          (FAN IN)

1. gen(nums ...int) — Generator
	ทำหน้าที่สร้าง Unbuffered Channel แล้วแตกเลข 2 และ 3 ออกมาเป็น Stream
	เมื่อส่งครบแล้วสั่ง close(out)

2. sq(in chan int) — Fan-Out (กระจายงาน)
	ใน main มีการเรียก sq(in) สองครั้ง ได้เป็น channel c1 และ c2
	กลไกสำคัญ: ทั้ง c1 และ c2 ต่างก็พยายามอ่านข้อมูลจาก channel in แผ่นเดียวกัน

	ตัวไหนแย่งอ่านได้ก่อนก็จะได้เลขนั้นไปยกกำลังสอง
	(เช่น Goroutine 1 อ่านได้ 2 ไปคิดเป็น 4, Goroutine 2 อ่านได้ 3 ไปคิดเป็น 9)

	เมื่อ in ถูกปิด ลูป for n := range in ของทั้งสองตัวจะจบลง
	และแยกกันปิด channel ของตัวเอง (c1, c2)


3. 	merge(cs ...chan int) — Fan-In (รวบรวมผลลัพธ์)
	ทำหน้าที่รับ channel หลายๆ ตัวเข้ามารวบให้เหลือ channel out เพียงอันเดียว
	ตั้ง WaitGroup: wg.Add(len(cs)) ตามจำนวน channel ที่ส่งเข้ามา (ในที่นี้คือ 2)
	รัน Goroutine ดึงข้อมูล: วนลูปสร้าง Goroutine แยกอ่านค่าจากแต่ละ channel แล้วโยนค่าเข้า out
		มีการส่ง (c) เข้าเป็นพารามิเตอร์ (ch chan int) เพื่อป้องกันปัญหาการแย่งใช้ตัวแปรลูป
		เมื่ออ่านจนจบ channel ไหน ให้เรียก wg.Done()
	Goroutine คอยปิด Channel: สร้างอีก Goroutine หนึ่งแยกไว้สั่ง wg.Wait()
	เพื่อรอให้ทุก worker ทำงานเสร็จก่อน แล้วค่อยสั่ง close(out) ป้องกันไม่ให้ main ติด Deadlock

4. main() — Consumer
	วนลูปอ่านค่าจาก merge(c1, c2) มาแสดงผล
	ผลลัพธ์จะเป็น 4 แล้วตามด้วย 9 หรือ 9 แล้วตามด้วย 4 ขึ้นอยู่กับว่า Goroutine
	ตัวไหนประมวลผลเสร็จและแย่งส่งเข้า merge ได้ก่อนกัน
*/
