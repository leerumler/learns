package main

import (
	"fmt"
)

type vertex struct {
	X int
	Y int
}

var (
	// Structs that are defined literally using the curly brace/comma syntax
	// are called struct literals.
	v1 = vertex{1, 2}

	// Struct literals also be listed using the Name: syntax.
	v2 = vertex{X: 1}

	// They can also be simply empty.
	v3 = vertex{}

	//
	p1 = &vertex{10, 10}
)

func lookAVertex() {

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

func literally() {

	fmt.Println(v1, v2, v3, p1)

}

func nowWithPointers(vertex *vertex) vertex {

	/* Stuff for main()
	   v := vertex{1, 2}
	   n := nowWithPointers(&v)
	   fmt.Println(n)
	*/
	pointer := *vertex

	pointer.X = 5

	return pointer

}

func main() {

	lookAVertex()

}
