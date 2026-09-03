package main

import "fmt"

func main() {
	var i interface{} = "hello"
	// A type assertion provides access to an interface value's underlying concrete value.
	// next statement asserts that the interface value i holds the concrete type string
	// and assigns the underlying type string value to the variable s
	s := i.(string)
	fmt.Printf("i.(string) ? : %s\n", s)

	/*
		To test whether an interface value holds a specific type,
		a type assertion can return two values:
			the underlying value and a boolean value that reports whether the assertion succeeded.
	*/
	s, ok := i.(string)
	fmt.Printf("s, ok := i.(string)  ? s: %v, \t ok: %v\n", s, ok)

	f, ok := i.(float64)
	fmt.Printf("s, ok := i.(float64) ? s: %v, \t ok: %v\n", s, ok)

	// next line will panic because i does not hold a float64, so the statement will trigger a panic.
	// f = i.(float64) // panic
	fmt.Println(f)
}
