package main

type MyType struct {
	i int
}

func (p *MyType) Get() int {
	return p.i
}

type MyInterface interface {
	Get() int
	Set(i int)
}

func (p *MyType) Set(i int) {
	p.i = i
}

func GetAndSet(x MyInterface) {}

func main() {
	var p MyType
	GetAndSet(&p)
}
