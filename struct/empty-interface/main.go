package main

import (
	"fmt"
)

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("describe interface :(%v, %T)\n", i, i)
}
