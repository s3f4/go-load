package library

import (
	"fmt"
	"runtime"
)

// WriteStack writes func caller stack
func WriteStack() {
	stackSlice := make([]byte, 512)
	s := runtime.Stack(stackSlice, false)
	fmt.Printf("\n%s", stackSlice[0:s])
	for c := 0; c < 5; c++ {
		fmt.Println(runtime.Caller(c))
	}
}
