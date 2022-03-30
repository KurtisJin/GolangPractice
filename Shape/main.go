package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{
		base:   2.0,
		height: 3.0,
	}
	squ := square{
		sideLength: 4.0,
	}

	printShape(tri)
	printShape(squ)
}

func printShape(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {

	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
