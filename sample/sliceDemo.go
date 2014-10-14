package main

import "fmt"

func main() {
	var s []int
	var a [10]int

	for i := 0; i < 10; i++ {
		a[i] = i
	}

	s = a[:] // s = a[0:len(a)]

	for i := 0; i < 10; i++ {
		fmt.Println(s[i])
	}
}
