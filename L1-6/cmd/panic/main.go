package main

import (
	"fmt"
	"time"
)

func main() {

	go gother()

	time.Sleep(5 * time.Second)
	fmt.Println("Программа завершена")

}

func gother() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	for i := 0; i < 3; i++ {
		fmt.Println("Выполнение программы...")
		time.Sleep(1 * time.Second)
	}
	panic("Горутина остановлена")

}
