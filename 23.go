package main

import (
	"fmt"
	"math/rand"
)

func Removefromslice(s []int, idx int) []int {
	if idx < 0 || idx >= len(s) {
		return s
	}
	return append(s[:idx], s[idx+1:]...)
}

func main() {
	s := make([]int, 10)
	for i := range s {
		s[i] = rand.Intn(100)
	}

	fmt.Println(s)
	s = Removefromslice(s, 0)
	fmt.Println(s)
}
