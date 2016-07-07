package pointers

import "fmt"

// DemonstratePointers briefly demonstrates how pointers work.
func DemonstratePointers() {
	fmt.Println("Creating two numbers...")
	first, second := 42, 2701

	//
	firstpointer := &first
	fmt.Println("Original value of first number:", first)
	fmt.Println("Original value of pointer to first number:", *firstpointer)

	//
	fmt.Println("Re-assign first pointer to 21.")
	*firstpointer = 21

	fmt.Println("New Value of first pointer:", first)
	fmt.Println("New Value of first:", first)

	//
	secondpointer := &second
	fmt.Println("Original value of second number:", second)
	fmt.Println("Original value of pointer to second number:", *secondpointer)

	fmt.Println("Divide pointer by 37.")
	*secondpointer = *secondpointer / 37

	fmt.Println("New value of second pointer:", *secondpointer)
	fmt.Println("New value of second number:", second)

}
