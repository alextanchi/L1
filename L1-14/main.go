package main

import (
	"fmt"
)

func main() {
	var x interface{} = false
	//задача решается с помощью специальной конструкции
	switch y := x.(type) {
	case int:
		fmt.Printf("%v это int", y)
	case string:
		fmt.Printf("%v это string", y)
	case bool:
		fmt.Printf("%v это bool", y)
	case chan int:
		fmt.Printf("%v это channel of int", y)

	}
}
