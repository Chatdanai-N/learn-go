package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {

		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {

		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

/*
Semaphore (เซมาฟอร์) คือ กลไกการควบคุมและจัดลำดับการเข้าถึงทรัพยากร (Synchronization Primitive)
 ในการเขียนโปรแกรมแบบประมวลผลพร้อมกัน (Concurrent Programming)
 เช่น Multi-threading หรือ Multi-processing

แนวคิดมาจาก "สัญญาณธงเซมาฟอร์" ที่ใช้ในทางเรือหรือรถไฟ
 เพื่อให้สัญญาณว่าเส้นทางนั้นว่างหรือมีรถไฟคันอื่นใช้อยู่หรือไม่


*/
