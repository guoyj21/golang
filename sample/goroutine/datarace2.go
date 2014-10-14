package main

import (
	"fmt"
)

func sharingIsCaring() {
	ch := make(chan int)
	go func() {
		n := 0
		n++
		ch <- n
	}()
	n := <-ch
	n++
	fmt.Println(n)
}

func main() {
	sharingIsCaring()
}
