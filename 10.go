package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := make([]float32, 10)
	for i := range s {
		s[i] = float32(rand.Intn(100)-50) + float32(rand.Intn(10))/10
	}
	fmt.Println(s, "\n")

	m := make(map[int][]float32)
	for _, v := range s {
		m[int(v)/10*10] = append(m[int(v)/10*10], v)
	}
	fmt.Println(m)
}
