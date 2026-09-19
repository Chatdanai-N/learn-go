package main

import "fmt"

func main() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter:", <-c3+<-c4)
}

func incrementor(s string) chan int {
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
จุดที่อาจเป็นปัญหา (หากผลลัพธ์ไม่ตรงตามที่คุณคาดหวัง)
หากคุณรู้สึกว่าโค้ดมีความผิดปกติ อาจเกิดจากจุดประสงค์ในการเขียนดังนี้ครับ:

สร้าง Variable s ไว้แต่ไม่ได้ใช้งาน: ในฟังก์ชัน incrementor(s string)
มีการรับพารามิเตอร์ s เข้ามา (เช่น "Foo:" และ "Bar:") แต่ข้างในฟังก์ชันไม่ได้ใช้ s ทำอะไรเลย
หากต้องการพิมพ์ออกมาดู ให้เพิ่ม fmt.Println(s, i) ในลูป

ไม่มี buffering ให้ channel: Channel ทั้งหมดเป็น unbuffered (make(chan int))
ซึ่งทำงานแบบบล็อก (Synchronous) ในเคสนี้ทำงานได้ดี แต่อาจบล็อกการทำงานทันทีหากไม่มีตัวรับ data


*/
