package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater than 10")
	} else {
		return fmt.Sprint("x is less than 10")
	}
}

/*
change go-lint to golangci-lint

golangci-lint --version
วิธีตรวจสอบผลลัพธ์: หากติดตั้งถูกต้อง ระบบจะแสดงเวอร์ชันของ golangci-lint ออกมา

golangci-lint run
วิธีตรวจสอบผลลัพธ์: ระบบจะสแกนโค้ด Go ทั้งหมดใน Directory ปัจจุบัน และแสดงข้อผิดพลาดหรือข้อเสนอแนะในการแก้ไขโค้ดออกมา


*/
