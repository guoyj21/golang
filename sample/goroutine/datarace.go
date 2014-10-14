package main

import (
	"fmt"
)

func main() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n++
		close(wait)
	}()
	n++
	<-wait
	fmt.Println(n)
}
