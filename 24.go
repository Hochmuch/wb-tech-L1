package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	var a, b Point
	a = NewPoint(0, 0)
	b = NewPoint(1, 1)
	fmt.Println(a.Distance(b))
}
