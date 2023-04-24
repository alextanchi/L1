package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go gother()
	time.Sleep(5 * time.Second)
	fmt.Println("Программа завершена")
}
func gother() {
	for {
		for i := 0; i < 3; i++ {
			fmt.Println("Выполнение программы...")
			time.Sleep(1 * time.Second)
		}
		runtime.Goexit() // останавливаем с помощью данной функции
	}

}
