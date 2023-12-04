package main

import (
	"fmt"
	"math/big"
)

func validate(s string, v *big.Int) bool {
	comparer := big.NewInt(2 << 20)
	fmt.Scanln(&s)
	if _, ok := v.SetString(s, 10); !ok {
		return false
	}
	if v := v.Cmp(comparer); v < 0 {
		return false
	}
	return true
}

func main() {
	a := new(big.Int)
	b := new(big.Int)
	v1 := ""
	v2 := ""
	fmt.Printf("please enter v1: ")
	if !validate(v1, a) {
		fmt.Println("Wrond data")
		return
	}
	fmt.Printf("please enter v2: ")
	if !validate(v2, b) {
		fmt.Println("Wrond data")
		return
	}

	res := new(big.Int)
	res.Sub(a, b)
	fmt.Printf("sub res is %v\n", res)

	res.Add(a, b)
	fmt.Printf("add res is %v\n", res)

	res.Div(a, b)
	zero := new(big.Int)
	zero.SetString("0", 10)
	if t := res.Cmp(zero); t == 0 {
		fmt.Println("div by zero")
	} else {
		fmt.Printf("div res is %v\n", res)
	}

	res.Mul(a, b)
	fmt.Printf("mult res is %v\n", res)
}
