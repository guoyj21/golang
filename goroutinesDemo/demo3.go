package main

var a string

var ch = make(chan int)

func f() {
	a = "Hello world"
	<-ch
}

func main() {
	go f()
	ch <- 0
	println(a)
}
