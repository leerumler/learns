package main

import (
	"fmt"
)

func whatsASlice() {

	// A slice is a dynamically-sized, flexible view in to the elements
	// of an underlying array.  They are declared just like an array, except
	// they don't have a size defined.
	var emptyslice []int
	fmt.Println(emptyslice)

	// There are a couple ways to make a slice, but the easiest is to
	// start with an array and take a pre-defined chunk of it.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	primeslice := primes[1:4]
	fmt.Println(primeslice)

}

func referencesToArrays() {
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

	// As such, when a slice is modified, the underlying array is modified,
	// and any other slices that shares the same underlying array will
	// reflect the change.

}

func main() {
	referencesToArrays()
}
