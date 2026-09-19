package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

/*
อธิบายลำดับการทำงานทีละขั้นตอน
1. เริ่มต้นทำงานที่ f()
	โปรแกรมลงทะเบียน defer พร้อม recover() เอาไว้ที่หน้าประตูของฟังก์ชัน f()
	พิมพ์ "Calling g." ออกมา
	เรียกใช้งานฟังก์ชัน g(0)

2. การทำงานช่วงขาไปของ g(i) (Recursion Stack)
โปรแกรมจะเรียก g() วนซ้ำไปเรื่อยๆ ตั้งแต่ i = 0 ถึง 3:

	g(0): ลงทะเบียน defer fmt.Println("Defer in g", 0) แล้วพิมพ์ "Printing in g 0" จากนั้นเรียก g(1)
	g(1): ลงทะเบียน defer fmt.Println("Defer in g", 1) แล้วพิมพ์ "Printing in g 1" จากนั้นเรียก g(2)
	g(2): ลงทะเบียน defer fmt.Println("Defer in g", 2) แล้วพิมพ์ "Printing in g 2" จากนั้นเรียก g(3)
	g(3): ลงทะเบียน defer fmt.Println("Defer in g", 3) แล้วพิมพ์ "Printing in g 3" จากนั้นเรียก g(4)

3. จุดเปลี่ยน: เกิด panic ที่ g(4)
	เมื่อเข้าสู่ g(4) เงื่อนไข if i > 3 กลายเป็นจริง:
	1.พิมพ์ "Panicking!"
	2.สั่ง panic("4") ทำให้โปรแกรม หยุดการทำงานปกติทัน
	บรรทัด fmt.Println("Returned normally from g.") ใน f() จะ ไม่มีทางถูกทำงาน

4.การทำงานช่วงขากลับ: การคลี่คาย Stack และ defer (Unwinding)
เมื่อเกิด panic ระบบจะถอยหลังออกจาก Call Stack และไล่รันคำสั่ง defer
ทั้งหมดที่สะสมไว้ในฟังก์ชัน g() ตามลำดับ LIFO (เข้าทีหลัง-ออกก่อน):
	1.รัน defer ของ g(3) ➔ พิมพ์ "Defer in g 3"
	2.รัน defer ของ g(2) ➔ พิมพ์ "Defer in g 2"
	3.รัน defer ของ g(1) ➔ พิมพ์ "Defer in g 1"
	4.รัน defer ของ g(0) ➔ พิมพ์ "Defer in g 0"

5. การกู้คืนระบบที่ f() ด้วย recover
	1.หลังจากคลี่คาย defer ของฟังก์ชัน g ทั้งหมดแล้ว สภาวะ panic จะลามถอยกลับมาถึงฟังก์ชัน f():
	2.ฟังก์ชัน f() เรียกใช้ defer ที่ดักรอไว้ตั้งแต่เริ่มต้น
	3.ฟังก์ชัน recover() ดักจับค่า panic (ซึ่งก็คือ "4") ได้สำเร็จ ทำให้ สภาวะ panic ยุติลง ณ จุดนี้
	4.ฟังก์ชัน f() จบการทำงานลงอย่างเรียบร้อย (ไม่ Crash)
	5.ควบคุมการทำงานกลับมาที่ main() และพิมพ์ "Returned normally from f." เป็นบรรทัดสุดท้าย
*/
