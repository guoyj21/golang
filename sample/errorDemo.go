package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("filename.txt")
	if err != nil {
		log.Fatal(err)
	}

}
