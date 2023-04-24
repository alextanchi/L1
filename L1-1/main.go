package main

import "fmt"

func main() {
	n := Action{
		&Human{"Alex"},
		"fast",
	}
	n.Run(n.Rate) // вызов метода Run из структуры Human
}

type Human struct {
	Name string
}

//  Метод Run для структуры Human

func (h *Human) Run(r string) {
	fmt.Printf("%s is running %s \n", h.Name, r)
}

// встраивание структуры Human в структуру Action
//(после этого появляется возможность использовать метод Run для структуры Action)

type Action struct {
	*Human
	Rate string
}
