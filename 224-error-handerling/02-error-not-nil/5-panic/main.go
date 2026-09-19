package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-flie.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		//log.Println("err happened", err)
		//log.Fatal(err)
		panic(err)
	}
}

/*
พฤติกรรมเมื่อเกิด Error
พิมพ์ Error + Stack Trace แล้วพังโปรแกรม (แต่ยังทำ defer ก่อนปิด)


เหมาะสำหรับ
Bug ร้ายแรงในระดับ Internal/Developer Error


หยุดการทำงานตามปกติของโปรแกรมทันทีเมื่อเจอ Error ร้ายแรง โดยระบบจะพิมพ์ข้อมูลความผิดพลาด
พร้อมกับ Call Stack Trace (สายการเรียกฟังก์ชันทั้งหมด)
ออกมาให้ดูว่า Error เกิดขึ้นที่ไฟล์ไหนและบรรทัดไหน

พฤติกรรมของ panic เมื่อถูกเรียกใช้งาน
หยุดทำงานทันที: ฟังก์ชันปัจจุบันจะหยุดทำงานกลางคัน
รัน defer ทั้งหมด: ระบบจะย้อนกลับไปทำงานในบล็อกคำสั่ง defer
ของฟังก์ชันปัจจุบันและฟังก์ชันก่อนหน้าที่เรียกมาทั้งหมด (ต่างจาก log.Fatal ที่ข้าม defer ไปเลย)
โปรแกรมพัง (Crash): หากไม่มีการใช้ recover()
ดักไว้ โปรแกรมจะจบการทำงานพร้อมพ่น Stack Trace ออกมา
*/
