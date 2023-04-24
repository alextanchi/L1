package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 5, 8, 4} // при идеальных условиях этот срез должен быть длиннее
	arr2 := []int{3, 4, 5, 8}
	fmt.Println(intersection(arr1, arr2))
}
func intersection(arr1, arr2 []int) []int {
	m := make(map[int]bool) // Карта для хранения элементов множества
	result := []int{}       // Срез для хранения пересечения множеств

	for _, v := range arr1 { // Добавляем элементы первого множества в карту
		m[v] = true
	}
	// Проверяем элементы второго множества на наличие в карте
	// и добавляем их в срез, если они присутствуют в обоих множествах
	for _, v := range arr2 {
		if m[v] {
			result = append(result, v)
		}
	}

	return result
}
