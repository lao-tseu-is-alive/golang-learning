package math_operation

import (
	"fmt"
	"testing"
)

func TestSub(t *testing.T) {
	numbers := []int{1, 2, 3}
	expected := -6
	actual := Sub(numbers)

	if actual != expected {
		t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
	}
}

func ExampleSub() {
	numbers := []int{5, 5}
	fmt.Println(Sub(numbers))
	// Output:
	// -10
}
