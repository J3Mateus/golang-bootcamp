package main

import (
	"fmt"
)

type Produto struct {
	nome  string
	preco float64
}

func main() {
	var precoTotal float64 = 0

	produtos := make([]Produto, 0)
	produtos = append(produtos, Produto{nome: "Camiseta", preco: 29.9})
	produtos = append(produtos, Produto{nome: "Calça", preco: 99.9})
	produtos = append(produtos, Produto{nome: "Tênis", preco: 199.9})

	for _, produto := range produtos {
		precoTotal += produto.preco
	}

	fmt.Printf("O preço total dos produtos é: R$ %.2f\n", precoTotal)
}
