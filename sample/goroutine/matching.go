package main

import (
	"fmt"
	"sync"
)

func Seek(name string, match chan string, waitGroup *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		//wait for someone to receive my message
	}

	waitGroup.Done()
}

func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1)

	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(people))

	for _, name := range people {
		go Seek(name, match, waitGroup)
	}

	waitGroup.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %s's message.\n", name)
	default:
		// There was no pending send operation
	}
}
