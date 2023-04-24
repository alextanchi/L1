package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	words := strings.Split(str, " ") // разбиваем строку на слова, получаем массив

	rev := make([]string, len(words)) //создаем массив куда будем "складывать" слова в обратном порядке

	for i, j := 0, len(words)-1; i <= len(words)-1; i, j = i+1, j-1 {
		rev[i] = words[j] // заполняем массив rev
	}

	fmt.Printf(strings.Join(rev, " ")) // превращаем массив слов в одну строку
}
