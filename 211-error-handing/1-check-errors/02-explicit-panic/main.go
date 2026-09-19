package main

import "fmt"

func main() {
	age := -5
	if age < 0 {
		panic("อายุติดลบไม่ได้!") // สั่ง panic พร้อมส่งข้อความแจ้งเตือน
	}
	fmt.Println("Age:", age)
}
