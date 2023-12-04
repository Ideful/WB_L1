package main

import (
	"fmt"
	"unicode"
)

func UniqueChecker(s string) bool {
	m := make(map[rune]struct{}, len(s))

	for _, v := range s {
		if unicode.IsLetter(v) {
			v = unicode.ToLower(v)
		}
		if _, c := m[v]; c == true {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}

func main() {
	s := "qwertyuiopasa"
	fmt.Println(UniqueChecker(s))
	s = "Qweasd"
	fmt.Println(UniqueChecker(s))
	s = "Asdfghzxc"
	fmt.Println(UniqueChecker(s))
	s = "qwertyuiopas"
	fmt.Println(UniqueChecker(s))
	s = "asdWDz"
	fmt.Println(UniqueChecker(s))

}
