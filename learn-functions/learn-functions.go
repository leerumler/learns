package main

import "fmt"

// from learn-addition.go
func add(x, y int) int {
	return x + y
}

// from learn-swap-strings.go
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 133))

	a, b := swap("hello", "world")
	fmt.Println(b, a)
}
