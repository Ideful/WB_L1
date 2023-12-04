package main

import "fmt"

func changebit(num *int64, bitpos, val int64) bool {
	if bitpos < 0 || bitpos > 63 || (val != 0 && val != 1) {
		return false
	}

	op := int64(1) << bitpos
	if val == 0 {
		*num &^= op
	} else { // val == 1
		*num |= op
	}
	return true
}

func main() {
	var num int64
	var bit_pos, val int64

	fmt.Println("enter your nubmer")
	fmt.Scan(&num)

	fmt.Println("enter bit_pos [0,63]")
	fmt.Scan(&bit_pos)

	fmt.Println("enter 0/1")
	fmt.Scan(&val)

	if ok := changebit(&num, bit_pos, val); ok == false {
		fmt.Print("wrong input")
		return
	}
	fmt.Printf("\n%d\n", num)
}
