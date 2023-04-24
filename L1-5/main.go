package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// Записываем данные в канал
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("Сообщение %d", i)
			<-time.After(600 * time.Millisecond)
		}
	}()

	timeout := time.After(5 * time.Second) //устанавливаем таймер

	for i := 0; ; i++ {
		select {
		case gopher := <-ch: // Получаем сообщение
			fmt.Println(gopher)
		case <-timeout: // Ждем окончания времени
			fmt.Println("Время вышло, программа остановлена")
			return
		}

	}

}
