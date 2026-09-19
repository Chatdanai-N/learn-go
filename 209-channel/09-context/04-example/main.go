package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0

		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Microsecond * 200)
				fmt.Println("Working", n)

			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num goroutines 2:", runtime.NumGoroutine())

	fmt.Println("abount to cancle context")
	cancel() // ส่งสัญญาณยกเลิกไปยัง ctx.Done()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num goroutines 3:", runtime.NumGoroutine())
}

/*

	[Main Thread]                             [Background Goroutine]
      │                                             │
      ├─► context.WithCancel()                      │
      ├─► Print (Goroutines: 1)                     │
      │                                             │
      ├─► go func() ────────(รันใน Background)────►┤
      │                                             │  วนลูปทำงานทุก 200 µs
      ├─► Sleep 2 วินาที                            │  ("Working 1", "Working 2"...)
      │                                             │
      ├─► Print (Goroutines: 2)                     │
      │                                             │
      ├─► cancel() ───(ส่งสัญญาณ Cancel)──────────►│  เจอ case <-ctx.Done():
      │                                             │  แล้วสั่ง return (จบการทำงาน!)
      ├─► Sleep 2 วินาที (รอ Goroutine จบจริง)       │
      │                                             │
      └─► Print (Goroutines: 1, Err: Canceled)      x ( Goroutine ตายแล้ว )


เมื่อเราเปิด Goroutine ทำงานแบบ Infinite Loop (for { ... }) ใน Background เราต้องมีกลไกบอก
ให้มันหยุดทำงานอย่างปลอดภัย (Clean Shutdown) เพื่อป้องกัน Goroutine Leak
ใน Go จะใช้ context.WithCancel ซึ่งจะคืนค่ามา 2 ตัว:
 1.ctx (Context Object): ส่งเข้าไปให้ Goroutine ย่อยคอยเช็กสัญญาณเตือนผ่าน <-ctx.Done()
 2.cancel (Cancel Function): ฟังก์ชันที่ฝั่ง Main จะเรียกใช้เมื่อต้องการ "กดปุ่มยกเลิก"

โค้ดนี้แสดงรูปแบบการใช้ Non-blocking Select (select + default) ร่วมกับ Context
เพื่อทำ Background Task ที่เปิดโอกาสให้ Cancel งานกลางคันได้อย่างสมบูรณ์แบบครับ

*/
