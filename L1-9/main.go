package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int)
	out := make(chan int)
	arr := [...]int{5, 6, 7, 8, 9, 10}
	go conveyer(in, out) //запускаем конвеер

	go func() { //Запускаем горутину записи данных в канал in
		for _, i := range arr {
			in <- i
		}
		close(in)
	}()

	time.Sleep(2 * time.Second)
	for z := range out { // выводим значения из канало out
		fmt.Println(z)
	}
}
func conveyer(in, out chan int) {
	for value := range in {
		out <- value * 2
	}
	close(out)

}
