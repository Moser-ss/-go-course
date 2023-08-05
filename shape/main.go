package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func main() {
	sq := square{sideLength: 2}
	t := triangle{base: 4, height: 8}

	printArea(sq)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println("The area is :" + fmt.Sprintf("%f", s.getArea()))
}
