package main

import (
	"03-funcoes-e-modulos/utils"
	"fmt"
)

func main() {

	num := 3
	resultado := utils.Even(num)

	if resultado == 1 {
		fmt.Println("É par!")
	} else {
		fmt.Println("É impar!")
	}

	num2 := 29
	if utils.IsPrime(num2) {
		fmt.Println("É primo!")
	} else {
		fmt.Println("Não é primo!")
	}
}
