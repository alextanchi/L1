package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	arr := []int{2, 4, 6, 8, 10}
	sum := 0

	for _, element := range arr {
		wg.Add(1)
		go SumOfSquares(element, &wg, &sum)
	}
	wg.Wait()
	fmt.Println("Сумма квадратов чисел:", sum)
}

func SumOfSquares(i int, wg *sync.WaitGroup, sum *int) {
	defer wg.Done()
	*sum += i * i
	fmt.Println("горутина с аргументом", i, "отработала")
}
