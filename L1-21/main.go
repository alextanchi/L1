package main

import "fmt"

//Есть структура Square с методом Area(), который вычисляет площадь квадрата.
//Нужно использовать этот метод в функции, которая ожидает интерфейс Shape с методом GetArea()

//Создаем структуру Square с методом Area(), который вычисляет площадь квадрата.
//Затем мы создаем интерфейс Shape с методом GetArea(), который мы будем использовать в функции printArea().

// Square - квадрат
type Square struct {
	side float64
}

// Area - вычисление площади квадрата
func (s *Square) Area() float64 {
	return s.side * s.side
}

// Shape - интерфейс для фигур
type Shape interface {
	GetArea() float64
}

//Создаем адаптер SquareAdapter, который адаптирует Square к интерфейсу Shape.
//В SquareAdapter мы используем композицию, чтобы хранить ссылку на Square.

// SquareAdapter - адаптер для квадрата
type SquareAdapter struct {
	square *Square
}

// GetArea - вычисление площади квадрата через адаптер
func (sa *SquareAdapter) GetArea() float64 {
	return sa.square.Area()
}

// Функция, которая ожидает интерфейс Shape
func printArea(shape Shape) {
	fmt.Println("Площадь фигуры:", shape.GetArea())
}

// Создаем экземпляр Square и SquareAdapter и передаем SquareAdapter в функцию printArea(),
// которая ожидает интерфейс Shape.
func main() {
	square := &Square{side: 5}
	squareAdapter := &SquareAdapter{square: square}
	printArea(squareAdapter)
}
