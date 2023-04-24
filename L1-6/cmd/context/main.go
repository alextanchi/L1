package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go gother(ctx)
	time.Sleep(3 * time.Second)
	cancel() //останавливем с помощью вызова функции cancel
	time.Sleep(5 * time.Second)
	fmt.Println("Программа завершена")
}

func gother(ctx context.Context) {
	for {
		select {
		default:
			fmt.Println("Выполнение программы...")
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			fmt.Println("Горутина остановлена")
			return
		}
	}
}
