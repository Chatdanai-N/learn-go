package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//f, err := os.Create("log.txt")

	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	// defer f.Close(): สั่งปิดไฟล์เมื่อฟังก์ชัน main ทำงานเสร็จ
	log.SetOutput(f)

	//log.SetOutput(f): นี่คือจุดสำคัญ เปลี่ยนปลายทางของแพ็กเกจ log
	// จากเดิมที่ปกติจะพิมพ์ออกทางหน้าจอ (Standard Error)
	// ให้เปลี่ยนไปเขียนบันทึกข้อมูลลงไฟล์ f (log.txt) แทน

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	} else {
		defer f2.Close()
	}

	fmt.Println("Check the log.txt file int the directory")

}
