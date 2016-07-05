package main

import (
	"fmt"
)

func whatsAnArray() {

	// Arrays are fixed-length in go, so they can't be resized.
	// The number in the brackets defines how longh the array is.
	var array [2]string

	// The array index starts at 0 and contains the number of elements
	// defined in the bracket.
	array[0] = "Hello"
	array[1] = "World"

	fmt.Println(array[0], array[1])

	// There is also a short syntax for declaring arrays using
	// curly braces and commas, much like structs.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func main() {
	whatsAnArray()
}
