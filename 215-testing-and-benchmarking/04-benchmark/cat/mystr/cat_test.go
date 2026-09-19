package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Cat(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func ExampleCat() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	// Output:
	// Shaken not stirred
}

func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output:
	// Shaken not stirred
}

const s = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}

/*
goos: windows
goarch: amd64
pkg: mymodule/215-testing-and-benchmarking/04-benchmark/cat/mystr
cpu: AMD Ryzen 7 9800X3D 8-Core Processor
BenchmarkCat-16            89494             12612 ns/op
BenchmarkJoin-16         2801684               442.4 ns/op


1. เปรียบเทียบความเร็วต่อการทำงาน 1 รอบ (ns/op)
BenchmarkJoin: ใช้เวลา 442.4 ns/op (442.4 นาโนวินาที ต่อรอบ)
BenchmarkCat: ใช้เวลา 12,612 ns/op (12,612 นาโนวินาที ต่อรอบ)

2. เปรียบเทียบจำนวนรอบที่รันได้ใน 1 วินาที
BenchmarkJoin: ทำงานได้ถึง 2,801,684 รอบ
BenchmarkCat: ทำงานได้เพียง 89,494 รอบ


ทำไม Join ถึงเร็วกว่า Cat มากขนาดนี้?

1. Cat (ใช้ + ต่อ String หรือ fmt.Sprintf ในลูป):

String ในภาษา Go เป็นประเภทข้อมูลแบบ Immutable (แก้ไขค่าเดิมไม่ได้)

ทุกครั้งที่มีการสั่งต่อ String ด้วย + ในลูป Go จำเป็นต้อง
สร้าง String ก้อนใหม่และจองหน่วยความจำใหม่ (Memory Reallocation)
รวมทั้งก๊อปปี้ข้อมูลเก่าใส่ข้อมูลใหม่เรื่อยๆ ซึ่งกินทรัพยากร CPU และความเร็วสูงมาก

2. Join (ใช้ strings.Join หรือ strings.Builder):
strings.Join จะทำการคำนวณความยาวรวมของข้อความทั้งหมดล่วงหน้าก่อน

จากนั้นจะทำการ จองพื้นที่หน่วยความจำเพียงครั้งเดียว (Single Allocation)
แล้วก๊อปปี้ข้อมูลลงไปในทีเดียว ทำให้ไม่เสียเวลาสร้างวัตถุใหม่ซ้ำๆ และเร็วกว่าแบบเห็นได้ชัดครับ

*/
