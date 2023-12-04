package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	ch := make(chan int)
	timer := time.NewTimer(time.Duration(n) * time.Second)

	go func() {
		defer close(ch)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				fmt.Println("returning")
				return
			default:
				select {
				case ch <- rand.Intn(100):
				case <-timer.C:
					fmt.Println("returning")
					return
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case val, ok := <-ch:
				if !ok {
					fmt.Println("Channel closed. Exiting.")
					return
				}
				fmt.Printf("Value from ch is %d\n", val)
			}
		}
	}()

	<-timer.C
	fmt.Println("exit")
	time.Sleep(4 * time.Second)
}
