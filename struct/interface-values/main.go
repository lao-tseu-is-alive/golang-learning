package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	/*
		If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
		In some languages this would trigger a null pointer exception,
		but in Go it is common to write methods that gracefully handle being called with a nil receiver
		(as with the method M in this example.)
	*/
	if t == nil {
		fmt.Printf("INTERFACE M for T string value: %v\n", "<nil>")
	} else {
		fmt.Printf("INTERFACE M for T string value: %v\n", t.S)
	}
}

type F float64

func (f F) M() {
	fmt.Printf("INTERFACE M for F float value: %v\n", f)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("describe interface :(%v, %T)\n", i, i)
}
