package main

import "fmt"

func main() {

	c := incrementor()
	cSum := puller(c)

	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func puller(c chan int) chan int {

	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

/*
โค้ดชุดนี้ทำงานได้ถูกต้อง และเป็นตัวอย่างที่ยอดเยี่ยมของ Pipeline Pattern ในภาษา Go ครับ
(การส่งต่อข้อมูลเป็นทอดๆ ผ่าน Channel เหมือนสายพานการผลิต)

ผลลัพธ์ที่ได้จากการรันคือพิมพ์เลข 45 ออกมาทางหน้าจอ
1. incrementor() — ผู้สร้างข้อมูล (Generator Stage)
	ทำหน้าที่สร้าง Unbuffered Channel ชื่อ out ขึ้นมา
	รัน Goroutine เพื่อยิงตัวเลข 0 ถึง 9 เข้าไปใน out
	เมื่อส่งครบ 10 ตัวแล้วสั่ง close(out) เพื่อบอกว่าหมดข้อมูลแล้ว
	คืนค่า Channel out ออกไปให้ผู้เรียกใช้งาน
2. puller(c) — ผู้ประมวลผลข้อมูล (Processor Stage)
	รับ Channel c (ซึ่งก็คือ out จาก incrementor) เข้ามาเป็น Input
	สร้าง Channel out ของตัวเองขึ้นมาใหม่เพื่อเตรียมส่งผลลัพธ์กลับ
	รัน Goroutine วนลูป for n := range c เพื่อดึงตัวเลขทีละตัวมาบวกสะสมไว้ในตัวแปร sum ($0+1+2+...+9 = 45$)
	เมื่อ c ถูกสั่ง close ลูปจะจบลง puller จึงยิงผลรวม sum เข้าไปใน out แล้วทำการ close(out)
3. main() — ผู้รับผลลัพธ์ปลายทาง (Consumer Stage)
	ต่อสายพาน: เรียก incrementor() ได้ c มา แล้วส่งต่อให้ puller(c) ได้ cSum กลับมา
	วนลูป for n := range cSum เพื่อรอรับผลลัพธ์ปลายทาง
	พิมพ์ค่า 45 ออกมา และหลุดลูปปิดโปรแกรมเรียบร้อยเมื่อ cSum ถูกสั่ง close


จุดเด่นของแนวทางนี้
	1. Clean Code & Encapsulation: แต่ละฟังก์ชันรับผิดชอบหน้าที่ของตัวเองชัดเจน
	และจัดการการสร้าง/ปิด Channel ภายในตัวเอง (Generator สร้างและเป็นคนปิด Channel เสมอ)

	2. Non-blocking Initialization: ทั้ง incrementor และ puller คืนค่า Channel ออกมาทันที
	แล้วปล่อยให้งานประมวลผลทำเบื้องหลังผ่าน Background Goroutines
*/
