package variables

import "fmt"

func ZeroValues() {
	var speed int    // numeric type
	var heat float64 // numeric type
	var off bool
	var brand string

	fmt.Println(speed)
	fmt.Println(heat)
	fmt.Println(off)

	// I've used printf to print an empty string
	// EXERCISE: Use Println to see what happens
	fmt.Printf("%q\n", brand)
	//it says: fmt.Println call has possible formatting directive %qprintf
}
