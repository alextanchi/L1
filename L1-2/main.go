package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	arr := []int{2, 4, 6, 8, 10}
	for _, element := range arr {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			fmt.Println("Квадрат числа", i, "=", i*i)
			wg.Done()
		}(element, &wg)
	}
	defer wg.Wait()
}
