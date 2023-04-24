package main

import (
	"fmt"
	"math"
)

func main() {
	// Задаем координаты точек
	a := NewPoint(0, 0)
	b := NewPoint(4, 6)

	distance := a.Distance(b) // Вычисляем расстояние

	fmt.Printf("Расстояние между точками = %.5f", distance)
}

// Создаем структуру Point

type Point struct {
	x float64
	y float64
}

// Конструктор для создания новой точки

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
//AB = √(Xb - Xa)^2 + (Yb - Ya)^2 --- формула из математики

func (a *Point) Distance(b *Point) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	fmt.Println(dx, dy)
	return math.Sqrt(dx*dx + dy*dy)
}
