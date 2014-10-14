package main

import "fmt"

func main() {
	s0 := []int{1, 2}
	s1 := append(s0, 3)
	s2 := append(s1, 4, 5)
	s3 := append(s2, s0...) //add slice

	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
}
