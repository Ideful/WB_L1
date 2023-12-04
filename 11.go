// Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

func main() {
	s1 := make(map[int]int32)
	s2 := make(map[int]int32)

	for i := 0; i < 10; i++ {
		s1[i] = int32(2 * i)
		s2[6+i] = int32(4*i + 1)
	}
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := make(map[int]int)
	for i := range s1 {
		if _, ok := s2[i]; ok == true {
			s3[i] = 42
		}
	}
	fmt.Println(s3)
}
