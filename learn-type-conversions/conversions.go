package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y int = 3, 4
	fmt.Println("Value of X: ", x)
	fmt.Println("Value of Y: ", y)

	var xsquared int = (x * x)
	var ysquared = (y * y)
	fmt.Println("X Squared: ", xsquared)
	fmt.Println("Y Squared: ", ysquared)

	squares := (xsquared + ysquared)
	fmt.Println("Squares: ", squares)

	var sqrt float64 = math.Sqrt(float64(squares))
	var sqrtConvert int = int(sqrt)
	fmt.Println("Float64 Square Root: ", sqrt)
	fmt.Println("Int Square Root: ", sqrtConvert)
}
