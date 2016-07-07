package loops

import "fmt"

// CountToAThousand counts from 1 to 1000.
func CountToAThousand() {
	fmt.Println("Sum starts at 1")
	for sum := 1; sum < 1000; {
		sum++
		fmt.Println("Running total", sum)
	}
}
