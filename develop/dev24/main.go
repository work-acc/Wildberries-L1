package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type point struct {
	x, y float64
}

func NewPoint(x, y float64) point {
	return point{x, y}
}

func (p point) Distance(other point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)
	distance := point1.Distance(point2)

	fmt.Printf("Distance between points: %.2f\n", distance)
}
