package main

import "fmt"

func main() {

	// receive-only channel
	c := make(<-chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("------")
	fmt.Printf("%T\n", c)
}

/*
	c := make(<-chan int, 2)
	ลูกศรอยู่หน้า chan (<-chan int)
	แปลว่าเรากำลังสร้าง Receive-only Channel (Channel ที่ดึงข้อมูลออกได้อย่างเดียว ห้ามส่งเข้า)

	c <- 42 💥 เกิด Error ตรงนี้ทันที!
	เราพยายามจะ ส่งข้อมูลเข้า ไปใน Channel ที่ถูกจำกัดสิทธิ์ให้
	รับข้อมูลได้อย่างเดียว คอมไพเลอร์จึงฟ้องทันทีว่า:

	invalid operation: cannot send to receive-only type <-chan int
*/
