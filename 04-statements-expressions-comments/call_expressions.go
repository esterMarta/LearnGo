package expressions

import (
	"fmt"
	"runtime"
)

func CallExpressions() {
	fmt.Println(runtime.NumCPU() + 1)
}
