package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func stopUsingChan(stop <-chan bool) {
	for {
		select {
		default:
			// Выполнение работы в горутине
			fmt.Println("workin...1")
			time.Sleep(1 * time.Second)
		case <-stop:
			// Получен сигнал об остановке
			fmt.Println("bye 1")
			return
		}
	}
}

func stopUsingReturn() {
	fmt.Println("workin...2")
	time.Sleep(1 * time.Second)
	fmt.Println("bye 2")
	return
}

func stopUsingCtx(ctx context.Context, cancel context.CancelFunc) {
	fmt.Println("workin...3")
	time.Sleep(1 * time.Second)
	fmt.Println("bye 3")
	cancel()
}

func stopUsingWg(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("workin...4")
}

func main() {
	stop := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)

	go stopUsingChan(stop)
	go stopUsingReturn()
	go stopUsingCtx(ctx, cancel)
	go stopUsingWg(&wg)

	time.Sleep(2 * time.Second)
	stop <- true
	wg.Wait()
	fmt.Println("bye 4")
	fmt.Println("Программа завершена")
}
