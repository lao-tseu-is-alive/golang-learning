package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golang-learning/numbers/mymath"
	"github.com/lao-tseu-is-alive/golog"
	"math"
)

func compareNewtonMethod(x float64) {
	myResult, err := mymath.Sqrt(x)
	if err == nil {
		fmt.Printf("Sqrt(%v) approximation with Newton method : %v\n", x, myResult)
		fmt.Printf("Sqrt(%v) from the Go math package         : %v\n", x, math.Sqrt(x))
		fmt.Printf("Difference of Go math with Newton method : %.18f\n", math.Sqrt(x)-myResult)
	} else {
		golog.Err("NO NEGATIVE NUMBERS PLEASE err : %v", err)
	}
}

func main() {

	compareNewtonMethod(2.0)
	compareNewtonMethod(3.0)
	compareNewtonMethod(4.0)
	compareNewtonMethod(49.0)
	compareNewtonMethod(-2)

}
