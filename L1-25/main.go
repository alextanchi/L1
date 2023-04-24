package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) { //через заданнаое время функция закончит работу
	<-time.After(time.Duration(seconds) * time.Second) //возвращает канал  через указанное время.

}

func main() {
	fmt.Println("Начало")
	Sleep(5) //задаем время в секундах
	fmt.Println("Конец")
}
