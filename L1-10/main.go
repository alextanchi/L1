package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)

	for _, temperature := range temperatures {
		Key := int(temperature/10) * 10 // получаем необходимый ключ путем целочисленного деления,
		// а затем умножения на 10 (для удоволетворения условиям задачи)
		m[Key] = append(m[Key], temperature) //в первом случае m[key] это ключ, а внутри функции append это уже слайс
	}

	for k, temperatures := range m { //выводим результат
		fmt.Println("Шаг:", k, "Значения температуры:", temperatures)
	}
}
