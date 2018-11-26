// Checks if we should go our separate ways.
package journey

// Sum adds up all passed interger parameters.
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
