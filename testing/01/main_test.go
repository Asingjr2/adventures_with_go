// Run test with "go test" or "go test -v".
// Tested func should start with TEST...
// Testing file should end with "file name" then "_test".

package main

import (
	"testing"
)
// Testing function must start with word test.
func TestMySum(t *testing.T) {

	type test struct {
		data []int
		answer int
	}
	// Multiple test example.  Create slice of test examples.
	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{50, 50}, 100},
		test{[]int{10, -1}, 9},
		test{[]int{1, 1,1}, 3},
	}

	for _, v := range tests {
		// Unfurling data as function parameters.
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Looking for ", v.answer, "Got ", x, "istead")
		}
	}
	// Single test example
	first := mySum(2,3) 
	if first != 5 {
		t.Error("Test should equal 5, got ", first )
	}
}
