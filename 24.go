package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Point struct {
	x int
	y int
}

func NewPoint() *Point {
	return &Point{
		x: int(rand.Intn(100)),
		y: int(rand.Intn(100)),
	}
}

func (p *Point) FindDistance(p2 *Point) int {
	return int(math.Sqrt(math.Pow(float64(p.x)-float64(p2.x), 2) + (math.Pow(float64(p.y)-float64(p2.y), 2))))
}

func main() {
	p1 := NewPoint()
	p2 := NewPoint()
	fmt.Println(p1.x, p1.y, p2.x, p2.y)
	fmt.Println(p1.FindDistance(p2))
}
