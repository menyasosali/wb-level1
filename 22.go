package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(2)
	b := big.NewInt(1000)

	// Умножение
	proizv := new(big.Int).Mul(a, b)
	fmt.Println(proizv)

	// Деление
	otn := new(big.Int).Div(b, a)
	fmt.Println(otn)

	// Сложенние
	suma := new(big.Int).Add(a, b)
	fmt.Println(suma)

	// Вычитание
	razn := new(big.Int).Sub(b, a)
	fmt.Println(razn)

}
