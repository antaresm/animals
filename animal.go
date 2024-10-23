package main

type Animal struct {
	Color  string
	Weight int
}

func (a Animal) Say() {
	println("-")
}

func (a Animal) SayColor() {
	println(a.Color)
}

type Dog struct {
	Animal
}

func (d Dog) Say() {
	println("Waf!")
}

type Cat struct {
	Animal
}

func (c Cat) Say() {
	println("Meow!")
}

func (a Animal) SayWeight() {
	println(a.Weight)
}
