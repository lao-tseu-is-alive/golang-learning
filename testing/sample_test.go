package main

import (
	"strconv"
	"testing"
)

/*
The testing package of the standard library provides support for the code testing needs.
The test function needs to fulfill the name pattern, TestXXX.
By default, the test tool looks for the file named xxx_test.go.
Note that each test function takes the T pointer argument,
which provides the useful methods for test control.
By the T struct pointer, the status of the test could be set.
For instance, the methods Fail and FailNow, cause the test to fail.
With the help of the T struct pointer, the test could be skipped by calling Skip, Skipf, or SkipNow.

The interesting method of the T pointer is the method Helper.
By calling the method Helper, the current function is marked as the helper function,
and if the FailNow (Fatal) is called within this function,
then the test output points to the code line where the function is called within the test,
as can be seen in the preceding sample code.
*/

func TestSampleOne(t *testing.T) {
	expected := "11"
	result := strconv.Itoa(10)
	compare(expected, result, t)
}

func TestSampleTwo(t *testing.T) {
	expected := "11"
	result := strconv.Itoa(10)
	compareWithHelper(expected, result, t)
}

func TestSampleThree(t *testing.T) {
	expected := "10"
	result := strconv.Itoa(10)
	compare(expected, result, t)
}

func compareWithHelper(expected, result string, t *testing.T) {
	t.Helper()
	if expected != result {
		t.Fatalf("Expected result %v does not match result %v",
			expected, result)
	}
}

func compare(expected, result string, t *testing.T) {
	if expected != result {
		t.Fatalf("Fail: Expected result %v does not match result %v",
			expected, result)
	}
	t.Logf("OK: Expected result %v = %v",
		expected, result)
}
