package main

import (
	"fmt"
)

func Generator(number chan int) {
	for i := 2; ; i++ {
		number <- i
	}
}

func Filter(number chan int, remain chan int, prime int) {
	for {
		i := <-number
		if i%prime != 0 {
			remain <- i
		}
	}
}

func main() {
	number := make(chan int, 100)
	go Generator(number)

	for i := 0; i < 100; i++ {
		prime := <-number
		fmt.Println(prime)

		remain := make(chan int, 10)

		go Filter(number, remain, prime)

		number = remain
	}
}
