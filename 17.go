package main

import (
	"fmt"
	"math/rand"
)

func binarySearch(arr []int, target int) (int, bool) {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid, true
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0, false
}

func main() {
	slice := make([]int, rand.Intn(40)+20)
	for i := range slice {
		slice[i] = rand.Intn(15)
	}

	target := rand.Intn(15)
	fmt.Println(slice, target)

	result, ok := binarySearch(slice, target)
	if !ok {
		fmt.Println("nfound")
	} else {
		fmt.Println(result)
	}
}
