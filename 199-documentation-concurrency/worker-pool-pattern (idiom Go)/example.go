func handle(queue chan *Request) {
    // Goroutine แต่ละตัวจะวนลูปดึงงานจาก queue ไปทำเรื่อยๆ
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // 1. สร้าง Worker Goroutines ตามจำนวน MaxOutstanding มารอไว้
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // 2. ยืนรอสัญญาณเลิกงาน (เช่น สั่ง Shutdown Server)
}

ข้อดี: ประสิทธิภาพสูงที่สุด เพราะไม่มี Overhead ในการสร้าง/ทำลาย Goroutine ใหม่เลย 
และจำนวน Goroutine จะนิ่งอยู่ที่ MaxOutstanding ตัวเสมอ ไม่มีทาง RAM บวมแน่นอน