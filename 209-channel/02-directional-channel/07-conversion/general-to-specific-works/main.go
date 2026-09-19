package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// general to specific converts
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

}

/*
General to Specific Converts
การแปลง Channel แบบทั่วไป ไปเป็น Channel แบบจำกัดสิทธิ์

// general to specific converts
fmt.Println("-----")
fmt.Printf("c\t%T\n", (<-chan int)(c))
fmt.Printf("c\t%T\n", (chan<- int)(c))

(<-chan int)(c)
เป็นการใช้ Type Conversion เพื่อแปลง c (ที่เป็น chan int แบบสองทาง)
ให้กลายเป็น <-chan int (Receive-only)
ผลลัพธ์: Go ยอมให้แปลงได้ เพราะเป็นการ "ลดสิทธิ์" จากที่อ่าน+เขียนได้ ให้เหลือแค่อ่านอย่างเดียว

(chan<- int)(c)
เป็นการแปลง c ให้กลายเป็น chan<- int (Send-only)
ผลลัพธ์: Go ยอมให้แปลงได้เช่นกัน เพราะเป็นการ "ลดสิทธิ์" ให้เหลือแค่ส่งเข้าอย่างเดียว
*/
