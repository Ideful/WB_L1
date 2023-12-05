// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, v := range sl {
		if _, ok := set[v]; ok != true {
			set[v] = struct{}{}
		}
	}
	fmt.Println(len(set), set)
}
