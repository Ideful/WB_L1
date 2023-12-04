package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := make([]int, 20)
	in := make(chan int)
	out := make(chan int)

	for i := range s {
		s[i] = rand.Intn(101)
	}

	go func() {
		for _, v := range s {
			in <- v
		}
		close(in)
	}()

	go func() {
		for v := range in {
			out <- 2 * v
		}
		close(out)
	}()

	i := 0
	for v := range out {
		fmt.Println(i, v)
		i++
	}

}
