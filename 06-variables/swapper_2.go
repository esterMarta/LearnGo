package variables

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Swapper #2
//
//  1. Swap the values of `red` and `blue` variables
//
//  2. Print them
//
// EXPECTED OUTPUT
//  blue red
// ---------------------------------------------------------

func Swapper2() {
	// UNCOMMENT THE CODE BELOW:

	red, blue := "red", "blue"
	// ?

	red, blue = blue, red

	fmt.Println(red, blue)
}
