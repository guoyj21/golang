package main

import (
	"fmt"
	"github.com/guoyj21/stackImpl/stack"
)

func main() {
	//	var s stack.Stack
	var s *stack.Stack = new(stack.Stack)
	s.Push("world")
	s.Push("Hello, ")

	for s.Size() > 0 {
		fmt.Print(s.Pop())
	}

	fmt.Println()
}
