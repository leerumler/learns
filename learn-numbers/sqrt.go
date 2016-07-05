package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	var z float64 = 50
	for n := 0; n < 10; n++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
		fmt.Println("Value after", n+1, "runs through the equation:", z)
	}
	return z
}

func main() {
	fmt.Println(math.Sqrt(6))
	fmt.Println(sqrt(6))
}
