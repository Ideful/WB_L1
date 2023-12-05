/*
Реализовать конкурентную запись данных в map.
*/

package main

import (
	"fmt"
	"sync"
)

type Cmap struct { // делаем свой класс и встраиавем туда мьютекс для корректной работы
	m map[int]string
	sync.Mutex
}

func NewCmap() *Cmap {
	m := make(map[int]string)
	c := Cmap{
		m: m,
	}
	return &c
}

func (c *Cmap) insert(n int, s string) {
	c.Lock()
	defer c.Unlock()
	c.m[n] = s
}

func main() {
	m := NewCmap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ { // конкурентная запись в map
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.insert(i, fmt.Sprintf("key%d", i))
			fmt.Println(i, m.m[i])
		}(i)
	}

	wg.Wait()

	fmt.Println(m)
}
