package main

import "fmt"

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the eve channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	//close(e) // ปิด channel eve
	//close(o) // ปิด channel odd

	q <- 0 // ส่งเลข 0 เข้า quit channel
}

/*


โค้ดส่วนนี้มีปัญหาเกี่ยวกับ การพยายามดึงข้อมูลจาก Channel ที่ถูกปิดไปแล้ว (Closed Channels) ร่วมกับการใช้ select ครับ

เมื่อสั่งรัน โค้ดนี้จะไม่หยุดแค่การพิมพ์ quit
แต่จะเกิด Infinite Loop (ลูปไม่รู้จบ)
พิมพ์ข้อความ from the eve channel: 0 หรือ from the odd channel: 0
ออกมาซ้ำๆ ไม่ยอมหยุด จนโปรแกรมค้าง!

3. พฤติกรรมของ select กับ Closed Channel (ต้นเหตุของ Bug 💥)
	ในภาษา Go มีกฎสำคัญเกี่ยวกับ Channel ที่ถูกปิดไปแล้ว ดังนี้:

	การดึงข้อมูลจาก Channel ที่ถูก close() ไปแล้ว จะไม่บล็อก (ไม่รอ)
	แต่จะคืนค่าเป็น Zero Value (ของ int คือ 0) ออกมาทันทีตลอดเวลา!

	เมื่อ eve และ odd ถูกปิด ตัว select จะมองว่าทั้งช่อง e, o, และ q
	มีข้อมูลพร้อมให้ดึงทั้งหมดพร้อมๆ กัน!


	ตัว select ใน Go จะทำการ สุ่มเลือก (Random Choose) ว่าจะเข้าไปทำ case ไหนก่อน:

	ถ้าสุ่มได้ case v := <-q: โปรแกรมจะพิมพ์ quit แล้ว return จบทำงานได้ (ถ้ารอดได้ด้วยดวง)

	แต่ถ้าสุ่มไปโดน case v := <-e: หรือ case v := <-o: ก่อน
	มันจะดึงค่า 0 (Zero Value ของ Channel ที่ปิดแล้ว) มาพิมพ์
	และวนลูป for ต่อไปเรื่อยๆ ไม่ยอมหลุดไปถึง case v := <-q: เสียที
	เกิดเป็น Infinite Loop ทันที



*/
