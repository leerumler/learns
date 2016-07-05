package main

import "fmt"

func main() {
	fmt.Println("Sum starts at 1")
	for sum := 1; sum < 1000000; {
		sum++
		fmt.Println("Running total", sum)
	}
}
