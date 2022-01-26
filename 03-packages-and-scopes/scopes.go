package packagesnscopes

import "fmt"

const ok = true

func Scopes() {
	var hello = "Hello"

	// hello and ok are visible here
	fmt.Println(hello, ok)
}
