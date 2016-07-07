package maths

import (
	"fmt"
	"math"
)

// PrintPi prints Pi.
func PrintPi() {
	fmt.Println(math.Pi)
}

// from sqrt.go
func sqrt(x float64) float64 {
	var z float64 = 50
	for n := 0; n < 10; n++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
		fmt.Println("Value after", n+1, "runs through the equation:", z)
	}
	return z
}

// TakeSquares compares the accuracy of a square root taken using Newton's
// method and a square root taken by Go's built-in math.Sqrt().
func TakeSquares() {
	fmt.Println(math.Sqrt(6))
	fmt.Println(sqrt(6))
}
