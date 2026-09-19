package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < 10; j++ { // เปลี่ยนจาก i เป็น j
				c <- j
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

/*
โค้ดชุดนี้ทำงานได้ถูกต้อง และเป็นรูปแบบ Fan-In Pattern
(การให้หลาย Goroutines ส่งข้อมูลมารวมกันที่ Channel เดียว)
 โดยสามารถรันและแสดงผลตัวเลข 0 ถึง 9 ออกมาทั้งหมด 100 ค่า (10 Goroutines × 10 รอบ)
 ได้โดยไม่เกิด Deadlock

 การทำงานทีละส่วน
 1. การเตรียม Channel และตัวแปร
 c := make(chan int): Channel หลักสำหรับรับส่งตัวเลข

 done := make(chan bool): Channel สำหรับรับสัญญาณแจ้งเตือนเมื่อ Goroutine ทำงานเสร็จ
 n := 10: จำนวน Goroutines ที่จะสร้างขึ้นมาทำงาน

 2. สร้าง 10 Goroutines เพื่อยิงข้อมูลเข้า Channel c
 ลูปสร้าง 10 Goroutines ให้ทำงานพร้อมกัน
 แต่ละ Goroutine จะส่งเลข 0 ถึง 9 เข้าไปใน c
 เมื่อส่งครบ 10 ตัวแล้ว Goroutine นั้นๆ จะยิง done <- true เพื่อบอกว่าตัวเองทำงานเสร็จแล้ว

 3. Goroutine สำหรับปิด Channel (Done Collector)
 สร้าง Goroutine มา 1 ตัวแยกต่างหาก เพื่อทำหน้าที่วนลูปรับสัญญาณ <-done ให้ครบ 10 ครั้ง (ตามจำนวน n)
 เมื่อรับครบทั้ง 10 สัญญาณ แสดงว่า Goroutine ผู้ส่งทุกตัวทำงานเสร็จแล้ว จึงสั่ง close(c) เพื่อปิด Channel c ได้อย่างปลอดภัย

 4. Main Goroutine อ่านและแสดงผล
 ลูป for n := range c ใน Main Goroutine จะคอยดึงข้อมูลจาก c มาพิมพ์ออกทางหน้าจอเรื่อยๆ
 ลูปนี้จะจบลงโดยอัตโนมัติเมื่อ Channel c ถูกสั่ง close(c) จาก Goroutine ในข้อ 3

 ทำไมโค้ดนี้ถึงไม่เกิด Deadlock (ต่างจากข้อที่แล้ว)?
 1.Main Goroutine ทำหน้าที่รับข้อมูลจาก c ตลอดเวลา ทำให้ Goroutine ทั้ง 10 ตัวยิง c <- i ได้เรื่อยๆ ไม่ติด Block
 2.เมื่อ Goroutine ใดส่งครบ 10 ตัว จะยิง done <- true ซึ่งมี Goroutine คอยรับ <-done รออยู่อีกตัวหนึ่ง ทำให้สายการทำงานไหลต่อเนื่องกันได้ทั้งหมด

*/
