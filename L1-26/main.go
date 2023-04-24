package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "ABCaXYZ"
	sr := []rune(s)

	for i := 0; i < len(s); i++ {
		sr[i] = unicode.ToLower(sr[i]) //обновляем срез sr: заполняем символами строки нижнего регистра,
		// для последующего сравнения
	}
	//в двойном цикле сравниваем каждый элемент с остальными элементами,
	// если повторяющихся нет - печатаем true
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if sr[i] == sr[j] {
				fmt.Println("false")
				return
			}
		}
	}
	fmt.Println("true")
}
