package main

import (
	"fmt"
	"time"
)

var ch chan string = make(chan string)

func sendMail() {
	fmt.Println("Sending mail...")
	time.Sleep(5 * 1e9)
	ch <- "done"
	fmt.Println("I already sent the mail")
}

func main() {

	go sendMail()

	fmt.Println("Waiting mail")
	result := <-ch
	fmt.Println("Got the mail " + result)
}
