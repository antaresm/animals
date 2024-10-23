package main

func main() {
	println("Hello World")

	animal := Animal{}
	animal.Say()

	dog := Dog{}
	dog.Color = "red"
	dog.Weight = 10
	dog.Say()
	dog.SayColor()

	cat := Cat{}
	cat.Color = "black"
	cat.Say()
	cat.SayColor()

}
