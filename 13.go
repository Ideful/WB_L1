// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&b, &a)

	fmt.Printf("a is %d | b is %d\n\n", a, b)
	a, b = b, a
	fmt.Printf("a is %d | b is %d\n", a, b)
}
