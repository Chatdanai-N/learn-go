package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exits")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}

/*

			   ┌── Goroutine (งาน 0)   ──┐
               ├── Goroutine (งาน 1)   ──┤
[populate] ──> ├── Goroutine (งาน 2)   ──┼─> [fanOutIn] ──> c2 ──> [main print]
  (c1)         │   ...                   │    (Fan-In)
               └── Goroutine (งาน 99)  ──┘
                       (Fan-Out)


Fan-Out: คือ Pattern ที่เรา "กระจายงานจาก Channel เดียว
ออกไปให้ Goroutine หลายๆ ตัวช่วยกันทำขนานกัน"
เพื่อลดเวลาประมวลผล (คล้ายการเพิ่มพนักงานมาช่วยกันเคลียร์คิวงาน)

Fan-In: คือการ "รวบรวมผลลัพธ์จากพนักงานเหล่านั้น
 กลับเข้าสู่ Channel เดียว (c2)" เพื่อส่งต่อให้ผู้รับ


***จังหวะ Fan-Out (กระจายงาน)
วนลูปอ่านงานจาก c1 ด้วย for v := range c1

ทุกๆ 1 งานที่อ่านได้ จะแตก Goroutine ใหม่ทันที 1 ตัว (go func(v2 int)...)
งานทั้ง 100 ชิ้นจึงถูกแตกออกไปทำขนานกัน 100 Goroutine
พร้อมกัน ไม่ต้องรอคนก่อนหน้าทำเสร็จ!

****จังหวะ Fan-In (รวมผลลัพธ์)
Goroutine ทั้ง 100 ตัว เมื่อทำงานใน timeConsumingWork เสร็จ
จะแย่งกันส่งผลลัพธ์กลับเข้ามาที่ Channel c2

ใช้ wg.Wait() เพื่อรอให้ Goroutine ทั้ง 100 ตัว
ประมวลผลและส่งค่าลง c2 ครบทุกตัวก่อน

เมื่อเสร็จครบแล้ว จึงเรียก close(c2)
เพื่อบอกฝั่ง main() ว่าหมดงานแล้ว


ข้อดีของ Pattern นี้

ประหยัดเวลามาก (Concurrency): หากงาน 1 ชิ้นใช้เวลา $1$ วินาที
	ทำแบบ Sequential (ทีละงาน): ใช้เวลาทั้งหมด $100$ วินาที
	ทำแบบ Fan-Out: งานทั้ง 100 ชิ้นรันพร้อมกัน จะใช้เวลาเหลือเท่ากับ ตัวที่ช้าที่สุดเพียงตัวเดียว (~1 วินาที)

Order ไม่สำคัญ: ลำดับการพิมพ์ผลลัพธ์ใน main() จะไม่เรียง $0$ ถึง $99$ เพราะตัวไหนประมวลผล
เสร็จก่อน ก็จะส่งเข้า c2 และถูกพิมพ์ออกมาก่อนทันที

*/
