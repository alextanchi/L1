package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	go gother(stop)
	time.Sleep(3 * time.Second)
	stop <- true //останавливаем с помощью передачи в канал флага
	time.Sleep(5 * time.Second)
	fmt.Println("Программа завершена")
}
func gother(stop chan bool) {

	for {
		select {
		default:
			fmt.Println("Выполнение программы...")
			time.Sleep(1 * time.Second)
		case <-stop:
			fmt.Println("Горутина остановлена")
			return
		}
	}
}
