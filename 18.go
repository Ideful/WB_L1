/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(c *Counter) {
			i := 0
			defer wg.Done()
			for i < 1000 {
				c.Increment()
				i++
			}
		}(c)
	}

	wg.Wait()
	fmt.Println(c.count)
}
