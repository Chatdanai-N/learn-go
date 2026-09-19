var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // ขอโควตา (ถ้าเต็มจะบล็อก)
    process(r)  // ทำงานหนัก
    <-sem       // คืนโควตา
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req) // ⚠️ ปัญหาอยู่ตรงนี้!
    }
}


การทำงาน: ทุกครั้งที่มี Request เข้ามา Serve จะปั๊ม Goroutine ใหม่ขึ้นมาทันที แล้วค่อยไปติดล็อก sem <- 1 อยู่ข้างใน handle

ปัญหา: สมมติมี Request ทะลักเข้ามา 1,000,000 งาน แต่ MaxOutstanding คือ 100... 
ผลคือโปรแกรมจะสร้าง Goroutine ขึ้นมา 1 ล้านตัว! ถึงแม้จะมีแค่ 100 ตัวได้ทำงานจริงๆ 
ส่วนอีก 999,900 ตัวยืนค้างกิน Memory เล่นๆ จน RAM หมด (Goroutine Leak)