package main

import "fmt"

type Human struct {
	name string
}

func (h Human) Greet() {
	fmt.Println("Hi! My name is", h.name)
}

type Action struct {
	Human
}

func main() {
	var example Action = Action{Human{name: "User"}}
	example.Greet()
}
