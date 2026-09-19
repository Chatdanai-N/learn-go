package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-flie.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		//log.Println("err happened", err)
		log.Fatal(err)
		//panic(err)
	}
}

/*
พฤติกรรมเมื่อเกิด Error
พิมพ์ข้อความ + วันเวลา แล้วสั่งดับโปรแกรมทันที (os.Exit(1))

เหมาะสำหรับ
Error ร้ายแรงที่ไม่สามารถทำงานต่อได้ (เช่น โหลด Config ไม่ได้)

ข้อควรระวังสำคัญของ log.Fatal
คำสั่ง log.Fatal จะใช้ os.Exit(1) ซึ่งจะ
หยุดโปรแกรมทันทีโดยข้ามการทำงานของคำสั่ง defer ทั้งหมด ในโปรแกรม
ดังนั้นหากมีงานที่ต้องคืน Resource (เช่น ปิด Database Connection หรือปิดไฟล์)
ควรหลีกเลี่ยงการใช้ log.Fatal ในฟังก์ชันย่อยๆ และควรย้ายไปไว้ที่ฟังก์ชัน main() เท่านั้น


*/
