package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)
	for _, s := range arr {
		set[s] = true
	}
	//получаем множество, где каждая уникальная строка представлена в виде ключа
	for m, _ := range set {
		fmt.Println(m)
	}
}
