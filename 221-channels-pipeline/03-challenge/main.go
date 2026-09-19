package main

import "fmt"

func main() {

	in := gen(1, 2, 3, 4, 5, 6)
	f := factorial(in)
	for n := range f {
		fmt.Println(n)
	}
}

func gen(n ...int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= len(n); i++ {
			out <- n[i-1]
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- fact(v)
		}
		close(out)
	}()

	return out
}

func fact(n int) int {

	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*

Challenge #1
-- Change the code above to execute 100 factorial computations concurrently and in paralled.
-- Use the "pipeline" pattern to accomplish this
*/

/*
[gen]  ---> ส่งเลข 1 ถึง 6 ผ่าน channel ---> [factorial] ---> คำนวณ Factorial ---> [main] พิมพ์ผลลัพธ์

1. ฟังก์ชัน gen (Generator)
	ทำหน้าที่แปลงข้อมูล Slice ให้กลายเป็น Stream ของข้อมูลผ่าน Channel
	รับพารามิเตอร์แบบ Variadic n ...int (ในที่นี้คือ 1, 2, 3, 4, 5, 6)

	คืนค่าเป็น Read-only Channel (<-chan int)
	สร้าง Goroutine มาวนลูปส่งค่า 1 ถึง len(n) (ก็คือ 1 ถึง 6) เข้า out channel
	เมื่อส่งครบแล้วสั่ง close(out) เพื่อบอกตัวรับว่าหมดข้อมูลแล้ว

2. ฟังก์ชัน factorial (Stage ประมวลผล)
	ทำหน้าที่รับค่าจาก Pipeline ขั้นก่อนหน้า มาคำนวณ แล้วส่งต่อไปยัง Stage ถัดไป
	รับ Read-only Channel in เข้ามา และคืนค่าเป็น Read-only Channel out

	สร้าง Goroutine ดึงค่าจาก in ด้วย for v := range in (วนลูปอ่านเรื่อยๆ จนกว่า in จะโดน close)
	ส่งค่า v ไปคำนวณในฟังก์ชัน fact(v) แล้วส่งผลลัพธ์เข้า out channel
	เมื่อ in ปิดตัวลง ลูปจะจบลง แล้วสั่ง close(out) ตาม


3. ฟังก์ชัน fact (Helper Function)
	ฟังก์ชันคำนวณค่า Factorial แบบปกติ (n! = n x(n-1) x... x1)
	เช่น fact(3) = 3x2x1 = 6
    	fact(5) = 5x4x3x2x1 = 120

4. ฟังก์ชัน main (Driver/Consumer)
	ต่อ Pipeline: gen $\rightarrow$ factorial
	วนลูปอ่านผลลัพธ์ for n := range f และพิมพ์ออกทางหน้าจอ
*/
