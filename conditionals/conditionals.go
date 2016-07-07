package conditionals

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// PrintSquares takes the square root of two numbers, one positive,
// one negative, demonstrating an unexported function that uses an if
// statement to take the square root of postive or negative numbers.
func PrintSquares() {
	fmt.Println(sqrt(2), sqrt(-4))
}
