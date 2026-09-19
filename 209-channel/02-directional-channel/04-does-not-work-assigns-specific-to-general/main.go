package main

import "fmt"

func main() {

	c := make(chan int)
	cr := make(<-chan int) // receive-only
	cs := make(chan<- int) // send-only

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// specific to general doesn't assign
	c = cr
	c = cs
}

/*
"เราไม่สามารถนำ Channel ที่จำกัดสิทธิ์ (Specific) มา assign
ให้กับ Channel แบบทั่วไป (General / Bidirectional) ได้"

c := make(chan int)     // General (Bidirectional): อ่านได้ และ ส่งได้
cr := make(<-chan int)  // Specific (Receive-only): อ่านได้อย่างเดียว
cs := make(chan<- int)  // Specific (Send-only):    ส่งได้อย่างเดียว

fmt.Printf("c\t%T\n", c)  // ได้: chan int
fmt.Printf("cr\t%T\n", cr) // ได้: <-chan int
fmt.Printf("cs\t%T\n", cs) // ได้: chan<- int

// specific to general doesn't assign
c = cr // ❌ Error!
c = cs // ❌ Error!
*/
