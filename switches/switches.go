package switches

import (
	"fmt"
	"runtime"
	"time"
)

// WhichOS uses switches to print a message based on the host OS.
func WhichOS() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd, plan9, windoze...
		fmt.Printf("%s", os)
	}
}

// WhenSaturday prints how far away Saturday is in "human" language.
func WhenSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("Almost there.")
	default:
		fmt.Println("Too far away...")
	}
}

// TimeOfDayBasedGreeting will offer a different greeting based upon
// what time it is.
func TimeOfDayBasedGreeting() {

	//
	t := time.Now().Hour()

	//
	morn := "Good Morning!"
	aft := "Good Afternoon."
	eve := "Good Evening."

	//
	if t > 4 {
		switch {
		case t < 12:
			fmt.Println(morn)
		case t < 17:
			fmt.Println(aft)
		default:
			fmt.Println(eve)
		}

		//
	} else {
		fmt.Println(eve)

	}
}
