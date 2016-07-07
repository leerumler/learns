package random

import (
	"fmt"
	"math/rand"
)

// FaveNum prints out a "random" favorite number.
func FaveNum() {
	fmt.Println("My favorite number is", rand.Int())
}
