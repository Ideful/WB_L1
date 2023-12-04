/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от
родительской структуры Human (аналог наследования).
*/
package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, my name is %s\n", h.name)
}

func main() {
	h := Human{
		name: "John",
		age:  22,
	}
	h.SayHi()
	a := Action{Human: h}
	a.SayHi()
}
