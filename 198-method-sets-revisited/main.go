package main

import (
	"fmt"
	"math"
)

type circle struct {
	redius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {

	c := &circle{
		redius: 5,
	}

	info(c)

}
