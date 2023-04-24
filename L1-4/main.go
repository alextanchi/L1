package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	var n int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&n)

	// Запускаем воркеров
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, &wg, ch)
	}

	// Записываем данные в канал
	go func() {
		for i := 1; ; i++ {
			ch <- fmt.Sprintf("Сообщение %d\n", i)
			<-time.After(time.Second * 1)
		}
	}()

	// Обрабатка сигналов завершения программы
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		close(ch)
		fmt.Println("Нажата комбинация Ctrl+C. Закрытие канала...")

	}()

	// Ждем завершение работы всех воркеров
	wg.Wait()
	fmt.Println("Все воркеры остановлены")
}
func worker(id int, wg *sync.WaitGroup, ch <-chan string) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-ch:
			if !ok { //эта конструкция позволяет избежать паники после попытки чтения из закрытого канала
				fmt.Printf("Воркер %d: канал закрыт \n", id)
				return
			}
			fmt.Printf("Воркер %d: %s", id, data)
		}
	}
}
