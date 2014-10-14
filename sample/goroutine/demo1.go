package main

import (
	"fmt"
	"time"
)

func myroutine(ch chan int) {
	fmt.Println("enter myroutine")

	ch <- 1
	fmt.Println("1111111")
	ch <- 2
	fmt.Println("2222222")
	ch <- 3
	fmt.Println("3333333")

	time.Sleep(3 * time.Second)
}

func main() {
	ch := make(chan int, 3)
	go myroutine(ch)
	<-ch
	fmt.Println("aaaaaaa")
	time.Sleep(2 * time.Second)
	//<-ch
	//fmt.Println("bbbbbbb")
	//<-ch
	//ch <- 4
}
