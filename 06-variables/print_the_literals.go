package variables

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the literals
//
//  1. Print a few integer literals
//
//  2. Print a few float literals
//
//  3. Print true and false bool constants
//
//  4. Print your name using a string literal
//
//  5. Print a non-english sentence using a string literal
//
// ---------------------------------------------------------

func PrintTheLiterals() {
	//integer literals
	fmt.Println(
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	)

	//float literals
	fmt.Println(
		1., 1.1, 1.2, 1.3, 1.4, 1.5,
	)

	//bool constants
	fmt.Println(
		true, false,
	)

	//string literals
	fmt.Println(
		"Ester", "Marta", "Tambunan",
	)

	//non-english literals
	fmt.Println(
		"brother", "わたしは、あなたを愛しています",
	)
}
