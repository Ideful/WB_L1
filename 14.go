/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

package main

import (
	"fmt"
	"sync"
)

func do(i interface{}) {
	switch v := i.(type) {
	default:
		fmt.Printf("Type is %T\n", v)
	}
}

type Qwe struct {
	asd int
}

func main() {
	do(21)
	do("hello")
	do(map[string]int{})
	do(struct{}{})
	var wg *sync.WaitGroup
	do(wg)
	q := Qwe{}
	do(q)
	do(make(chan int))
	do(true)
}
