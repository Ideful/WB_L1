package main

import "fmt"

type Dog struct{}

func (dog *Dog) WoofWoof() {
	fmt.Println("Гав-Гав")
}

type Cat struct{}

func (cat *Cat) MeowMeow(isCall bool) {
	if isCall {
		fmt.Println("мяу-мяу")
	}
}

//

type AnimalAdapter interface { // создаем общий интерфейс
	Reaction()
}

type DogAdapter struct {
	*Dog
}

type CatAdapter struct {
	*Cat
}

func NewDogAdapter(dog *Dog) AnimalAdapter { // адаптер для собаки
	return &DogAdapter{dog}
}

func NewCatAdapter(cat *Cat) AnimalAdapter { // адаптер для кошки
	return &CatAdapter{cat}
}

func (adapter *DogAdapter) Reaction() {
	adapter.WoofWoof()
}

func (adapter *CatAdapter) Reaction() {
	adapter.MeowMeow(true)
}

func main() {
	d := Dog{}
	c := Cat{}
	d1 := NewDogAdapter(&d)
	c1 := NewCatAdapter(&c)
	// теперь оба объекта реализовывают общий интерфейс и могут использовать метод Reaction()
	d1.Reaction()
	c1.Reaction()
}
