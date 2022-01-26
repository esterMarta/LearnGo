package variables

import (
	"fmt"
	"path"
)

// ---------------------------------------------------------
// EXERCISE: Discard The File
//
//  1. Print only the directory using `path.Split`
//
//  2. Discard the file part
//
// RESTRICTION
//  Use short declaration
//
// EXPECTED OUTPUT
//  secret/
// ---------------------------------------------------------

func DiscardTheFile() string {
	// UNCOMMENT THE CODE BELOW:

	// ? ?= path.Split("secret/file.txt")
	dir, _ := path.Split("secret/test/file.txt")

	fmt.Println(dir)
	return ""
}
