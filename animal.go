package main

type Animal struct {
	Color  string
	Weight int
}

func (a Animal) Say() {
	println("-")
}
