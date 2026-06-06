package main

import "fmt"

const (
	// iota เริ่มต้นที่ 0 แต่ใช้ขีดล่าง (_) เพื่อข้ามค่านี้ไป
	_ = iota
	a
	b
	c
	d
	e
	f
)

func main() {

	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a) // 1*2^1 = 2 | 10
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b) // 1*2^2 = 4 | 100
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c) // 1*2^3 = 8  | 1000
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d) // 16
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e) // 32
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f) // 64
}
