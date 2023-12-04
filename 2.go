/*
Написать программу, которая конкурентно рассчитает значение
квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	a := [5]float64{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fmt.Println(math.Pow(a[idx], 2))
		}(i)
	}
	wg.Wait()
}
