/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var n int // переменная, к-ая будет отвечать за время работы программы в секундах
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	ch := make(chan int) // создаем небуферизированный канал и таймер и примитив синхронизации wg
	timer := time.NewTimer(time.Duration(n) * time.Second)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() { // функция посылает данные в канал и по истечении таймера выходит
		defer wg.Done()
		defer close(ch)
		for {
			select {
			case <-timer.C:
				fmt.Println("\treturning")
				return
			default:
				v := rand.Intn(40)
				fmt.Printf("sending %d to channel\n", v)
				ch <- v
			}
		}
	}()

	go func() { // функция получает данные из канала пока есть возможность и печатает данные. как перестает - выходит
		defer wg.Done()
		for {
			select {
			case val, ok := <-ch:
				if !ok {
					fmt.Println("\tChannel closed. Exiting.")
					return
				}
				fmt.Printf("Value from ch is %d\n", val)
			}
		}
	}()

	// <-timer.C
	wg.Wait()
	fmt.Println("\texit")
}
