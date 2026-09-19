func handle(queue chan *Request) {
	for req := range queue {
		// ประมวลผลฟังก์ชัน req.f(req.args)
		// แล้วยิงผลลัพธ์กลับไปที่ resultChan ของ Client ตัวนั้นๆ
		req.resultChan <- req.f(req.args)
	}
}

คอยดึง Request มาคำนวณ แล้วส่งผลลัพธ์ย้อนกลับไปทาง resultChan ที่แนบมากับ Request นั้นๆ