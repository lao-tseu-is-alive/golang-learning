package main

import (
	"fmt"
)

var integer int64 = 32500
var floatNum float64 = 22000.456
var floatNum2 = 123456.789

func main() {

	// Common way how to print the decimal
	// number
	fmt.Printf("%d \n", integer)

	// Always show the sign
	fmt.Printf("%+d \n", integer)

	// Print in other base X -16, o-8, b -2, d - 10
	fmt.Printf("%X \n", integer)
	fmt.Printf("%#X \n", integer)

	// Padding with leading zeros
	fmt.Printf("%010d \n", integer)

	// Left padding with spaces
	fmt.Printf("% 10d \n", integer)

	// Right padding
	fmt.Printf("% -10d \n", integer)

	// Print floating point number
	fmt.Printf("floating point number : %f \n", floatNum)
	fmt.Printf("floating point number : %f \n", floatNum2)

	// Floating-point number with limited precision = 5
	fmt.Printf("float with precision 5: %.5f \n", floatNum)
	fmt.Printf("float with precision 5: %.5f \n", floatNum2)

	// Floating-point number with limited precision = 5 & right padding
	fmt.Printf("float precision 5 align 10: %16.5f \n", floatNum)
	fmt.Printf("float precision 5 align 10: %16.5f \n", floatNum2)

	// Floating-point number
	// in scientific notation
	fmt.Printf("%e \n", floatNum)
	fmt.Printf("%e \n", floatNum2)

	// Floating-point number
	// %e for large exponents
	// or %f otherwise
	fmt.Printf("%g \n", floatNum)
	fmt.Printf("%g \n", floatNum2)

}
