/*
Дана последовательность температурных колебаний
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
*/

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

	m := make(map[int][]float32) // для группировки используем мапу
	for _, v := range s {
		m[int(v)/10*10] = append(m[int(v)/10*10], v) // для корректных диапазонов приводим к инту, потом делим и умножаем на 10
	}
	fmt.Println(m)
}
