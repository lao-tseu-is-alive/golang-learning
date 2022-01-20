package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

// getArea returns the area of the triangle
func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

type square struct {
	sideLength float64
}

// getArea returns the area of the square
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Printf("Area of shape is : %.3f\n", s.getArea())
}

func main() {
	myTriangle := triangle{
		base:   10,
		height: 6,
	}
	mySquare := square{sideLength: 10}
	printArea(myTriangle)
	printArea(mySquare)
}
