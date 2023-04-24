package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			fmt.Printf("Горутина %v отработала\n", i)
			m[i] = i + 1
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		if value, ok := m[i]; ok { //смотрим есть ли элемент в мапе, если есть печатаем
			fmt.Printf("Ключ: %v, значение: %v\n", i, value)
		}
	}

}
