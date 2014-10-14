package main

type MyType struct {
	i int
}

func (p *MyType) Get() int {
	return p.i
}

func main() {
	//var pm *MyType = new(MyType)
	var pm = new(MyType)
	pm.i = 100
	var n = pm.Get()
	println(n)
}
