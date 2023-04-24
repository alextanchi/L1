package main

import (
	"fmt"
	"math/big"
)

func main() {
	//Используем пакет big для работы с большими числами, а также методов которые определены для типа big.Int
	a := big.NewInt(1048576)    // 2^20
	b := big.NewInt(1073741824) // 2^30

	// Умножение
	c := new(big.Int).Mul(a, b)
	fmt.Printf("%s * %s = %s\n", a, b, c)

	// Деление
	d := new(big.Int).Div(b, a)
	fmt.Printf("%s / %s = %s\n", b, a, d)

	// Сложение
	e := new(big.Int).Add(a, b)
	fmt.Printf("%s + %s = %s\n", a, b, e)

	// Вычитание
	f := new(big.Int).Sub(b, a)
	fmt.Printf("%s - %s = %s\n", b, a, f)
}
