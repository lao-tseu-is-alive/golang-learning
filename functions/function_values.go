package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

/*
Functions are values too. They can be passed around just like other values.
Function values may be used as function arguments and return values.
*/
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	add := func(x, y float64) float64 {
		return (x + y)
	}
	fmt.Printf("add(5, 12) = %v \n", add(5, 12))
	fmt.Printf("compute(add) will call add(3,4) = %v \n", compute(add))

	fmt.Printf("compute(hypot) will call hypot(3,4) = %v \n", compute(hypot))
	fmt.Println(compute(math.Pow))
}
