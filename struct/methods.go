package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
	a method is just a function with a receiver argument.

here the VertexInt type
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
You can declare methods with pointer receivers.
This means the receiver type has the literal syntax *T for some type T.
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
