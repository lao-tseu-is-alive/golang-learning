package math_operation

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("[1, 2, 3, 4, 5]", testSumFunc([]int{1, 2, 3, 4, 5}, 15))
	t.Run("[1, 2, 3, 4, -5]", testSumFunc([]int{1, 2, 3, 4, -5}, 5))

}

func testSumFunc(numbers []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := Sum(numbers)
		if actual != expected {
			t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
		}
	}
}

func ExampleSum() {
	/* run your tests using go test -v . inside this directory
	... Go uses the output comments section at the bottom of an ExampleXxx() function
	to determine what the expected output is, and then when tests are run it compares
	the actual output with the expected output in the comments and will trigger
	a failed test if these don’t match.
	This makes it incredibly easy to test and write example code at the same time.

	On top of creating easy to follow test cases, examples are also used
	to generate examples that are displayed inside of generated documentation.
	For example, the example above can be used to generate docs for our math_operation package.
	you can check it by running godoc -http=:6060

	*/
	numbers := []int{5, 5, 5}
	fmt.Println(Sum(numbers))
	// Output:
	// 15
}

func TestDatabase(t *testing.T) {
	// Pretend to use the db
	fmt.Println(db.Url)
}

var db struct {
	Url string
}

func TestMain(m *testing.M) {
	// here we can initialize what we want before to run the tests
	// Pretend to open our DB connection
	db.Url = os.Getenv("DATABASE_URL")
	if db.Url == "" {
		db.Url = "localhost:5432"
	}

	flag.Parse()
	exitCode := m.Run() // here we run all the tests

	// now we can do some cleanup before exiting
	// Pretend to close our DB connection
	db.Url = ""

	// Exit
	os.Exit(exitCode)
}
