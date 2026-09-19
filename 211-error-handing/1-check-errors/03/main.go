package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	// การใส่ defer ไว้ข้างหน้า หมายถึง "เลื่อนการทำงานของ f.Close()
	// ไปรันในวินาทีสุดท้ายที่ฟังก์ชัน main ทำงานเสร็จ" ช่วยป้องกันปัญหาการลืมปิดไฟล์

	r := strings.NewReader("John Stone")
	io.Copy(f, r)

}
