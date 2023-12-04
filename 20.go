package main

import (
	"bytes"
	"fmt"
	"strings"
)

func revertwords(s string) string {
	sl := strings.Split(s, " ")
	var buffer bytes.Buffer
	l := len(sl)

	for i := l - 1; i >= 0; i-- {
		buffer.WriteString(sl[i])
		if i != 0 {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}

func main() {
	s := "snow dog sun tree"
	fmt.Println(s)
	fmt.Println(revertwords(s))
}
