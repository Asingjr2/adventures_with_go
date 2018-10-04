// Error testing in golang requires description of error type as explicit exceptions do not exist
package main

import (
	"testing"
)

func TestCountWord(t *testing.T) {
	// countword should return correct len of string
	if countWord("hey") != 3 {
		t.Errorf("functon is not working correctly")
	}
}
