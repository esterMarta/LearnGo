package expressions

import "fmt"

func ExecutionFlow() {
	fmt.Println("Hello!")

	// Statements change the execution flow
	// Especially the control flow statements like `if`
	if 5 > 1 {
		fmt.Println("bigger")
	}
}
