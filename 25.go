// Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

func OwnSleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	fmt.Println(0)
	OwnSleep(3 * time.Second)
	fmt.Println(1)
}
