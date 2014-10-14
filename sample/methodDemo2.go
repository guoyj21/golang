package main

type MyInt int

func (p MyInt) Get() int {
	return int(p)
}

func f(i int) {}

func main() {
	var v MyInt
	v = v * v
	f(int(v))
	//	f(v)
}
