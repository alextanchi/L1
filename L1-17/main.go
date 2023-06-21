package main

import (
	"fmt"
)

func main() {
	array := []int{5, 6, 7, 8, 9, 10, 12, 14}
	search := binarySearch(array, 14) // в конце пишем элемент который хотим найти, если элемента нет в массиве вернется -1
	fmt.Printf("%v", search)
}

func binarySearch(array []int, value int) int {
	low := 0
	high := len(array) - 1

	for low < high {
		mid := (low + high) / 2 //ищем "средний" элемент в массиве

		guess := array[mid]
		if guess == value {
			return mid
		} else if guess > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if array[low] == value {
		return low
	}
	return -1
}
