package types

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	toBe          = false
	maxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

// BasicTypes demonstrates some of the basic types in go.
func BasicTypes() {
	const f = "%T(%v)\n"
	fmt.Printf(f, toBe, toBe)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, z, z)

}

// Conversions demonstrates how to convert between types.
func Conversions() {
	var x, y int = 3, 4
	fmt.Println("Value of X: ", x)
	fmt.Println("Value of Y: ", y)

	var xsquared = (x * x)
	var ysquared = (y * y)
	fmt.Println("X Squared: ", xsquared)
	fmt.Println("Y Squared: ", ysquared)

	squares := (xsquared + ysquared)
	fmt.Println("Squares: ", squares)

	var sqrt = math.Sqrt(float64(squares))
	var sqrtConvert = int(sqrt)
	fmt.Println("Float64 Square Root: ", sqrt)
	fmt.Println("Int Square Root: ", sqrtConvert)

}
