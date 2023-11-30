package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// создаем массив, канал и синхронизатор
	a := [5]float64{2, 4, 6, 8, 10}
	c := make(chan float64, len(a))
	var wg sync.WaitGroup

	for i := 0; i < len(a); i++ {
		wg.Add(1)        // +1 к счетчику синхронизатора
		go func(i int) { // запускаем горутину
			defer wg.Done()        // -1 к счетчику синхронизатора
			c <- math.Pow(a[i], 2) // в канал пишем квадрат элемента массива
		}(i)
	}

	wg.Wait() // блокируем дальнейшее выполнение пока счетчик не будет 0
	close(c)  // закрываем канал

	var sum float64
	for val := range c {
		sum += val
	}
	fmt.Println(sum, 100+64+36+16+4)
}
