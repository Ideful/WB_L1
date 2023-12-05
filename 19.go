/*
Разработать программу, которая переворачивает подаваемую на ход строку 
(например: «главрыба — абырвалг»). 
Символы могут быть unicode.
*/

package main

import "fmt"

func revert(str string) string {
	rs := []rune(str)
	l := len(rs)

	for i := 0; i < l/2; i++ {
		rs[i], rs[l-i-1] = rs[l-i-1], rs[i]
	}

	return string(rs)
}

func main() {
	str := "Главрыба"
	str = revert(str)
	fmt.Println(str)
}
