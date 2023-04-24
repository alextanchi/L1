package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done() //уведомляем WaitGroup о завершении работы горутины
		for i := 0; i < 3; i++ {
			fmt.Println("Выполнение программы...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Горутина остановлена")
	}()

	wg.Wait()
	time.Sleep(5 * time.Second)
	fmt.Println("Программа завершена")
}
