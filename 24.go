package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// NewPoint - конструктор для создания новых точек
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// DistanceTo - функция для вычисления расстояния между точками
func (p *Point) DistanceTo(q *Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p := NewPoint(0, 0)
	q := NewPoint(3, 4)
	distance := p.DistanceTo(q)
	fmt.Printf("Расстояние между точками (%v,%v) и (%v,%v) равно %v\n",
		p.x, p.y, q.x, q.y, distance)
}
