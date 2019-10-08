/*
Code based on the article from  Jon Calhoun :
https://www.calhoun.io/how-to-test-with-go/
*/
// Package math_operation provides basic math operation to demonstrate testing and documenting in golang
package math_operation

// Sub subtract up all of numbers in a slice and returns the resulting operation.
func Sub(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum -= n
	}
	return sum
}
