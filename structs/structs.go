package structs

import (
	"fmt"
)

type vertex struct {
	X int
	Y int
}

// LookAVertex demonstrates structs by creating and moving a vertex.
func LookAVertex() {

	// Structs hold collections of fields, usually variables or other structs.
	// Structs can be created directly using a curly brace/comma syntax.
	fmt.Println("A vertex plots two points on a plane, X and Y.")
	v := vertex{1, 2}

	// Fields in structs can be accessed using a dot.
	fmt.Println("Original value of X:", v.X)
	fmt.Println("Original value of Y:", v.Y)

	// They can also be assigned using a dot.
	v.X = 4
	v.Y = 1

	fmt.Println("Look, the vertex is moving!")
	fmt.Println("New value of X:", v.X)
	fmt.Println("New value of Y:", v.Y)
}

// Literally demonstrates literally defining structs.
func Literally() {
	// Structs that are defined literally using the curly brace/comma syntax
	// are called struct literals.
	v1 := vertex{1, 2}

	// Struct literals also be listed using the Name: syntax.
	v2 := vertex{X: 1}

	// They can also be simply empty.
	v3 := vertex{}

	//
	p1 := &vertex{10, 10}
	fmt.Println(v1, v2, v3, p1)

}

func nowWithPointers(vertex *vertex) vertex {

	pointer := *vertex

	pointer.X = 5

	return pointer

}

// NowWithPointers demonstrates how to structs also with pointers!
func NowWithPointers() {
	v := vertex{1, 2}
	n := nowWithPointers(&v)
	fmt.Println(n)

}
