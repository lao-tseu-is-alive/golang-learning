package main

import (
	"fmt"
	"strconv"
	"testing"
)

var testData = []int{10, 11, 017}

/*
Creating subtests
In some cases, it is useful to create a set of tests that could have a similar setup or clean-up code.
This could be done without having a separate function for each test.

The T struct of the testing package also provides the Run method that could be used to run the nested tests.
The Run method requires the name of the subtest and the test function that will be executed.
This approach could be beneficial while using, for example, the table driven tests.
The code sample just uses a simple slice of int values as an input.

Execute the tests by go test -v.
*/

func TestSampleSubtest(t *testing.T) {
	expected := "10"
	for _, val := range testData {
		tc := val
		t.Run(fmt.Sprintf("input = %d", tc), func(t *testing.T) {
			if expected != strconv.Itoa(tc) {
				t.Fail()
			}
		})
	}
}
