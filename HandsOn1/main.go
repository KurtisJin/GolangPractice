package main

import (
	"fmt"
	"math"
)

type square struct {
	sideLength float64
}
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {

	sq := square{
		sideLength: 10.4,
	}
	cir := circle{
		6.2,
	}
	printArea(sq)
	printArea(cir)

}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func printArea(s shape) {
	fmt.Println(s.area())
}
