package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golang-learning/numbers/mymath"
	"github.com/lao-tseu-is-alive/golog"
)

func main() {
	val, err := mymath.Sqrt(2)
	if err != nil {
		golog.Err("NO NEGATIVE NUMBER PLEASE err : %v", err)
	}
	fmt.Printf("Sqrt(2) = %v\n", val)
	val, err = mymath.Sqrt(-2)
	if err != nil {
		golog.Err("NO NEGATIVE NUMBER PLEASE err : %v", err)
	}
	fmt.Printf("Sqrt(2) = %v\n", val)

}
