package saying

import "fmt"

func Greet(s string) string {

	return fmt.Sprint("Welcome my dear ", s)
}

/*
command benchmark
 go test -bench Greet

 output:
 goos: windows
 goarch: amd64
 pkg: mymodule/215-testing-and-benchmarking/04-benchmark/live/saying
 cpu: AMD Ryzen 7 9800X3D 8-Core Processor
 BenchmarkGreet-16       36217668                33.42 ns/op
 PASS
 ok      mymodule/215-testing-and-benchmarking/04-benchmark/live/saying  1.459s


1. ข้อมูลระบบ (Environment Info)
	goos: windows: รันอยู่บนระบบปฏิบัติการ Windows
	goarch: amd64: รันบนสถาปัตยกรรม CPU แบบ 64-bit
	pkg: mymodule/.../saying: พาธของ Package ที่ถูกนำมาทดสอบ
	cpu: AMD Ryzen 7 9800X3D ...: ชนิดของ CPU ที่ใช้ประมวลผลการทดสอบนี้

2. ผลการทดสอบ (Benchmark Results)
	บรรทัด BenchmarkGreet-16   36217668   33.42 ns/op มีความหมายดังนี้:
	GOMAXPROCS -16 จำนวน Logical CPU Core (Threads) ที่ Go ดึงมาใช้ประมวลผลระหว่างรันเทสต์
	จำนวนรอบที่รัน 36,217,668 Go รันวนลูปฟังก์ชัน Greet("James") ซ้ำไปทั้งหมด 36.2 ล้านรอบ ภายในเวลาประมาณ 1 วินาที เพื่อให้ได้ค่าเฉลี่ยที่แม่นยำที่สุด
	เวลาเฉลี่ยต่อรอบ 33.42 ns/op ฟังก์ชัน Greet ใช้เวลาทำงานเฉลี่ยรอบละ 33.42 นาโนวินาที (1 ns = 1 ในพันล้านวินาที)


Coverage
go test -coverprofile c.out
go tool coveer -html c.out
*/
