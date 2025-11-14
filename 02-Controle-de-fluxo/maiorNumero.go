package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Entre com 2 números: ")
	fmt.Scan(&num1, &num2)

	if num1 < num2 {
		println("O maior entre os digitos inserido é o número: ", num2)
	} else {
		println("O maior entre os digitos inserido é o número: ", num1)
	}

}
