package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero1, numero2, resultado float64
	var opcao int

	for {
		fmt.Println("\n=== Calculadora ===")
		fmt.Println("1 - Soma")
		fmt.Println("2 - Subtração ")
		fmt.Println("3 - Multiplicação")
		fmt.Println("4 - Divisão")
		fmt.Println("5 - Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Println("Entre com 2 números: ")
			fmt.Scan(&numero1, &numero2)
			resultado = soma(numero1, numero2)
			fmt.Println("Resultado: ", resultado)

		case 2:
			fmt.Println("Entre com 2 números: ")
			fmt.Scan(&numero1, &numero2)
			resultado = subtração(numero1, numero2)
			fmt.Println("Resultado: ", resultado)

		case 3:
			fmt.Println("Entre com 2 números: ")
			fmt.Scan(&numero1, &numero2)
			resultado = multiplica(numero1, numero2)
			fmt.Println("Resultado: ", resultado)

		case 4:
			fmt.Println("Entre com 2 números: ")
			fmt.Scan(&numero1, &numero2)
			resultado, err := divisao(numero1, numero2)

			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Resultado:", resultado)
			}

		case 5:
			fmt.Print("encerrando a calculadora..!\n")
			return
		default:
			fmt.Print("opção invalida, digite novamente!\n")
		}

	}

}

func soma(n1 float64, n2 float64) float64 {
	return n1 + n2
}

func subtração(n1 float64, n2 float64) float64 {
	return n1 - n2
}

func multiplica(n1 float64, n2 float64) float64 {
	return n1 * n2
}

func divisao(n1 float64, n2 float64) (float64, error) {
	if n2 == 0 {
		return 0, errors.New("Não é possivel dividir por zero")

	} else {
		return n1 / n2, nil
	}

}
