package functions

import "fmt"

// from learn-addition.go
func add(x, y int) int {
	return x + y
}

// DoAddition uses an unexported function to add two numbers.
func DoAddition() {
	fmt.Println(add(42, 133))
}

// from learn-swap-strings.go
func swap(x, y string) (string, string) {
	return y, x
}

// StringSwap uses an unexported function to swap two strings.
func StringSwap() {
	a, b := swap("hello", "world")
	fmt.Println(b, a)
}
