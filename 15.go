package main

var justString string

// в данной реалзиации мы используем ненужную нам паять, так как juststring ссылается на данные из v
func someFunc() {
	v := createHugeString(1 << 5)
	justString = v[:100]
}

// лучше вот так, т.к. с помощью ф-ции fn возвращается копия
func someFunc2() {
	v := createHugeString(1 << 10)
	justString = fn(v)
}

func fn(s string) string {
	return s[:100]
}

func main() {
	someFunc()
	someFunc2()
}
