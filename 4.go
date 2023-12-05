/*

Реализовать постоянную запись данных в канал (главный поток). 
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. 
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. 
Выбрать и обосновать способ завершения работы всех воркеров.
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	// "sync"
	// "time"
)

func main() {
	// делаем канал и определяем кол-во воркеров + ошибка
	ch := make(chan int)
	var c int
	fmt.Println("enter workers amount")
	if _, err := fmt.Scan(&c); err != nil {
		return
	}

	// создаем контекст и триггер отмены используя контекст.
	ctx, cancel := context.WithCancel(context.Background())
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		<-signalCh
		fmt.Println("Terminating")
		cancel()
	}()

	go func() { // publisher
		for {
			select {
			case <-ctx.Done(): // отмена контекста
				fmt.Println("pub stops")
				close(ch)
				return
			default:
				ch <- rand.Intn(100)
			}
		}
	}()

	var wg sync.WaitGroup // чтобы завершить горутины receiver'a

	for i := 0; i < c; i++ {
		wg.Add(1)
		go func(i int) { // receiver
			defer func(wg *sync.WaitGroup) { // закрывает все горутины
				wg.Done()
				fmt.Printf("worker %d terminated\n", i)
			}(&wg)
			for {
				select {
				case val, ok := <-ch: // в канале есть данные
					if !ok {
						return
					}
					fmt.Printf("worker %d\t get %d\n", i, val)
				case <-ctx.Done(): // отмена контекста
					return
				}
			}
		}(i)
	}

	wg.Wait()
}
