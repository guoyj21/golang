package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(1000)
		fmt.Println("print this message in goroutine")
	}()

	fmt.Println("I am outside goroutine")
}
