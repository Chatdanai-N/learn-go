package main

import "fmt"

func main() {

	// send-only channel
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("------")
	fmt.Printf("%T\n", c)
}

/*
การใส่เครื่องหมาย <- ไว้หลังคำว่า chan แบบนี้ (chan<- int)
เป็นการประกาศให้ c เป็น Send-only Channel (Channel ที่ส่งข้อมูลเข้าได้อย่างเดียว ห้ามอ่านออก)

	c := make(chan<- int, 2)
	สร้าง Channel แบบ Send-only ที่มี Buffer ขนาด 2 ช่อง

	c <- 42 และ c <- 43
	ส่วนนี้ทำงานได้ปกติ เพราะ Channel นี้อนุญาตให้ส่งเข้าได้

	fmt.Println(<-c) 💥 เกิด Error ตรงนี้!
	พยายามจะอ่านข้อมูลออกจาก c (<-c) แต่ Go ไม่อนุญาตให้อ่านข้อมูลจาก Send-only Channel คอมไพเลอร์จะฟ้องทันที:

	invalid operation: cannot receive from send-only channel c (variable of type chan<- int)
*/
