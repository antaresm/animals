package main

type Animal struct {
	Color  string
	Weight int
}

func (a Animal) Say() {
	println("-")
}

type Dog struct {
	Animal
}

func (d Dog) Say() {
	println("Waf!")
}
