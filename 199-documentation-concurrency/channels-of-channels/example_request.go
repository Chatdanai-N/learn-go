type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int // <--- ช่องทางส่งผลลัพธ์กลับ
}

// Client สร้าง Request พร้อมสร้าง Channel มารอรับคำตอบ
request := &Request{[]int{3, 4, 5}, sum, make(chan int)}

// ส่ง Request เข้าไปในคิวหลักของ Server
clientRequests <- request

// ยืนรอรับคำตอบจากช่องทางส่วนตัวของตัวเอง
fmt.Printf("answer: %d\n", <-request.resultChan)