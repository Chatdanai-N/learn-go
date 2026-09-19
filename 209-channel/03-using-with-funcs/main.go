package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	go bar(c)

	fmt.Println("abount to exit")
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}

/*
จุดสังเกตสำคัญ: แม้ c จะเป็นแบบ Bidirectional (chan int)
แต่เมื่อถูกส่งเข้าไปใน foo และ bar ตัว Go จะ ลดสิทธิ์ (Convert)
ให้โดยอัตโนมัติ:
	foo(c chan<- int) -> ถูกจำกัดให้ ส่งได้อย่างเดียว
	bar(c <-chan int) -> ถูกจำกัดให้ รับได้อย่างเดียวรูปแบบนี้ช่วยการันตีความปลอดภัย (Type Safety) ว่าฟังก์ชันจะไม่ทำงานผิดหน้าที่

	สิ่งที่ต้องระวัง: การรัน Goroutine ใน main ต้องระวังเรื่องโปรแกรมชัตดาวน์ตัดหน้า
	(ต้องจัดการเรื่อง synchronization ให้ดี)

*/
