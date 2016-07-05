package main

import (
	"fmt"
)

func whatsASlice() {

	// A slice is a dynamically-sized, flexible view in to the elements
	// of an underlying array.  They are declared just like an array, except
	// they don't have a size defined.
	var emptyslice []int
	fmt.Println("This is an empty slice", emptyslice)

	// There are a couple ways to make a slice, but the easiest is to
	// start with an array and take a pre-defined chunk of it.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("An array full of primes:", primes)
	primeslice := primes[1:4]
	fmt.Println("Slice taken from an array full of primes:", primeslice)

	// A slice literal is declared just like an array literal,
	// without the length.
	arraylit := [4]int{0, 2, 7, 11}
	fmt.Println("Array literal:", arraylit)
	slicelit := []int{1, 3, 8, 12}
	fmt.Println("Slice literal:", slicelit)

	// Unlike arrays, slices have both a length and a capacity.
	// The length is the number of elements it contains, and  the
	// capacity is the length of the underlying array.

	//

}

func refsToArrays() {
	// A slice doesn't store any data; it describes a section of an
	// underlying array.  It's a "view" in to the array's contents.
	topfive := [5]string{
		"Shailene Woodley",
		"Natalie Dormer",
		"Zac Efron",
		"Sophie Turner",
		"Scarlett Johansson",
	}
	fmt.Println("Celebrity top five:", topfive)

	// If we take two slices from the same underlying array, they will both
	// continue to reference that array.  The data isn't copied: the memory
	// is shared.
	topthree := topfive[0:3]
	lowthree := topfive[2:5]
	fmt.Println("Top three:", topthree)
	fmt.Println("Bottom three:", lowthree)

	// When a slice is modified, the underlying array is modified, and any
	// slices that share that underlying array will reflect the change.
	lowthree[0] = "Johnny Depp"
	fmt.Println("Top five:", topfive)
	fmt.Println("Top three:", topthree)
	fmt.Println("Bottom three:", lowthree)

}

func howToSlice() {
	// When slicing an array, we can omit the high or low bounds and accept
	// the defaults:  zero for the low bound and the length of the array.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Array full of primes:", primes)

	firstthree := primes[:3]
	fmt.Println("First three primes:", firstthree)
	skipthree := primes[3:]
	fmt.Println("Skip first three:", skipthree)

	// The same can be done to another slice, because you can in fact
	// slice another slice.
	primeslice := primes[1:4]
	fmt.Println("Slice full of primes:", primeslice)

	firsttwo := primeslice[:2]
	fmt.Println("First two from slice:", firsttwo)
	skiptwo := primeslice[2:]
	fmt.Println("Skip first two from slice:", skiptwo)
}

func main() {
	howToSlice()
}
