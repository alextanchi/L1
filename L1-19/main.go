package main

import (
	"fmt"
)

func main() {
	s := "главрыба"
	sr := []rune(s) // создаем срез рун

	for i, j := 0, len(sr)-1; i < j; i, j = i+1, j-1 {
		sr[i], sr[j] = sr[j], sr[i] //меняем элементы местами
	}

	fmt.Println(string(sr))

}
