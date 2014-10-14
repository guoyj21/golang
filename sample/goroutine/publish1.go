package main

import (
	"fmt"
	"time"
)

func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("Breaking News:", text)
	}()
}

func main() {
	Publish("A goroutine starts a new thread of execution", 5*time.Second)
	fmt.Println("Let's hope the news will published before I leave")

	// wait for the news to be published
	time.Sleep(10 * time.Second)

	fmt.Println("10 seconds later: I'm leaving now.")
}
